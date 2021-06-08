package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"service-transaction/src"
	"service-transaction/src/config"
	"strings"
	"time"
)

func NewService(repo src.RepositoryList) src.ServiceList {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) Transaction(data []byte) (int, string, error) {

	status, ts, err := ConvertJsonToTransction(data)
	if err != nil {
		return status, "", err
	}

	ts.TransferTime, ts.EndDateReversal = SetData()

	status, err = s.CheckPassword(ts.Payer, ts.Password)
	if err != nil {
		return status, "", err
	}

	status, err = s.CheckUserType(ts.Payer)
	if err != nil {
		return status, "", err
	}

	status, nextBalance, err := s.CheckSufficientBalance(ts.Payer, ts.Value)
	if err != nil {
		return status, "", err
	}

	status, err = s.CheckApiResponse()
	if err != nil {
		return status, "", err
	}

	status, transferID, err := s.ExecuteTransaction(ts.Value, nextBalance, ts.Payer, ts.Payee, 1, ts.TransferTime, ts.EndDateReversal)
	if err != nil {
		return status, "", err
	}

	status, err = s.SendNotification(transferID)
	if err != nil {
		return status, "", err
	}

	return status, `{"message":"transação efetuada com sucesso"}`, nil
}

func (s *Service) CheckPassword(id int, pass string) (int, error) {

	var query strings.Builder
	var count int

	fmt.Fprintf(&query, `SELECT COUNT(id) FROM mydb.users where id = %d and password = '%s';`, id, pass)

	rows, err := s.Repository.Select(query.String())
	if err != nil {
		log.Println(fmt.Errorf("CheckPassword Select(): %w", err))
		return http.StatusInternalServerError, errors.New("erro no banco de dados")
	}

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Println(fmt.Errorf("CheckPassword rows.Scan(): %w", err))
			return http.StatusInternalServerError, errors.New("erro no banco de dados")
		}
		if count < 1 {
			return http.StatusUnauthorized, errors.New("senha incorreta")
		}
	}

	return http.StatusOK, nil
}

func (s *Service) CheckUserType(id int) (int, error) {

	var query strings.Builder
	var types int

	fmt.Fprintf(&query, `SELECT type FROM mydb.users where id = %d;`, id)

	rows, err := s.Repository.Select(query.String())
	if err != nil {
		log.Println(fmt.Errorf("checkUserType Select(): %w", err))
		return http.StatusInternalServerError, errors.New("erro no banco de dados")
	}

	for rows.Next() {
		err := rows.Scan(&types)
		if err != nil {
			log.Println(fmt.Errorf("CheckUserType rows.Scan(): %w", err))
			return http.StatusInternalServerError, errors.New("erro no banco de dados")
		}
		if types != 1 {
			return http.StatusForbidden, errors.New("permissão negada para realizar transferências")
		}
	}

	return http.StatusOK, nil
}

func (s *Service) CheckSufficientBalance(id int, value float64) (int, float64, error) {

	var query strings.Builder
	var balance, nextBalance float64

	fmt.Fprintf(&query, `SELECT balance FROM mydb.users where id = %d;`, id)

	rows, err := s.Repository.Select(query.String())
	if err != nil {
		log.Println(fmt.Errorf("CheckSufficientBalance Select(): %w", err))
		return http.StatusInternalServerError, 0, errors.New("erro no banco de dados")
	}

	for rows.Next() {
		err := rows.Scan(&balance)
		if err != nil {
			log.Println(fmt.Errorf("CheckSufficientBalance rows.Scan(): %w", err))
			return http.StatusInternalServerError, 0, errors.New("erro no banco de dados")
		}
		nextBalance = balance - value
		if nextBalance < 0 {
			return http.StatusForbidden, 0, errors.New("saldo insuficiente")
		}
	}

	return http.StatusOK, nextBalance, nil
}

func (s *Service) CheckApiResponse() (int, error) {
	client := http.Client{}
	data := &struct {
		Message string `json:"message"`
	}{}

	req, err := http.NewRequest(http.MethodGet, config.C.GetString("http.api_check"), nil)
	if err != nil {
		log.Println("CheckApiResponse NewRequest() %w", err)
		return http.StatusInternalServerError, errors.New("erro da requisição externa")
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		log.Println("CheckApiResponse NewRequest() %w", err)
		return http.StatusInternalServerError, errors.New("erro da requisição externa")
	}

	got, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("CheckApiResponse ReadAll() %w", err)
		return http.StatusInternalServerError, errors.New("erro interno")
	}

	err = json.Unmarshal(got, data)
	if err != nil {
		log.Println("CheckApiResponse Unmarshal() %w", err)
		return http.StatusInternalServerError, errors.New("erro interno")
	}

	if data.Message != "Autorizado" {
		return http.StatusUnauthorized, errors.New("não autorizado")
	}

	return http.StatusOK, nil
}

func (s *Service) ExecuteTransaction(value, nextBalance float64, payer, payee, types int, transferTime, endDateReversal time.Time) (int, int, error) {

	var res sql.Result
	var err error
	var transferID int64
	MySQLTimeFormat := "2006-01-02 15:04:05"
	tt := transferTime.Format(MySQLTimeFormat)
	ed := endDateReversal.Format(MySQLTimeFormat)

	var query strings.Builder
	if types == 1 {
		fmt.Fprintf(&query, `INSERT INTO mydb.transaction (value, payer, payee, end_reversal_date, createdate, updatedate) VALUES (%f, %d, %d, '%s', '%s', '%s');`,
			value, payer, payee, ed, tt, tt)

		res, err = s.Repository.Insert(query.String())
		if err != nil {
			log.Println(fmt.Errorf("ExecuteTransaction Select(): %w", err))
			return http.StatusInternalServerError, 0, errors.New("erro no banco de dados")
		}
	}

	query.Reset()

	fmt.Fprintf(&query, `UPDATE mydb.users SET balance = balance - %f WHERE (id = %d);`, value, payer)

	_, err = s.Repository.Insert(query.String())
	if err != nil {
		log.Println(fmt.Errorf("ExecuteTransaction Insert(): %w", err))
		return http.StatusInternalServerError, 0, errors.New("erro no banco de dados")
	}

	query.Reset()

	fmt.Fprintf(&query, `UPDATE mydb.users SET balance = balance + %f WHERE (id = %d);`, value, payee)

	_, err = s.Repository.Insert(query.String())
	if err != nil {
		log.Println(fmt.Errorf("ExecuteTransaction Insert(): %w", err))
		return http.StatusInternalServerError, 0, errors.New("erro no banco de dados")
	}
	if types == 1 {
		transferID, err = res.LastInsertId()
		if err != nil {
			log.Println(fmt.Errorf("ExecuteTransaction LastInsertId(): %w", err))
			return http.StatusInternalServerError, 0, errors.New("erro no banco de dados")
		}
	}
	return http.StatusOK, int(transferID), nil
}

func (s *Service) SendNotification(transferID int) (int, error) {
	client := http.Client{}
	send := &Notification{}

	var query strings.Builder

	fmt.Fprintf(&query, `select t.value, t.email_payee, t.payee_id, t.payee_name, s.email_payer, s.payer_id, 
		s.payer_name, t.transfer_time from (select ts.value, us.email as email_payee, 
		ts.payee as payee_id, us.name as payee_name , ts.createdate as transfer_time
		from mydb.transaction ts join mydb.users us on ts.payee = us.id where ts.id = %d) t
		join (select us.email as email_payer, ts.payer as payer_id, us.name as payer_name from mydb.transaction ts 
		join mydb.users us on ts.payer = us.id where ts.id = %d) s;`, transferID, transferID)

	rows, err := s.Repository.Select(query.String())
	if err != nil {
		log.Println(fmt.Errorf("SendNotification Select(): %w", err))
		return http.StatusInternalServerError, errors.New("erro no banco de dados")
	}

	for rows.Next() {
		err := rows.Scan(&send.Value, &send.Email_payee, &send.Payee_id, &send.Payee_name, &send.Email_payer, &send.Payer_id, &send.Payer_name, &send.Transfer_time)
		if err != nil {
			log.Println(fmt.Errorf("SendNotification rows.Scan(): %w", err))
			return http.StatusInternalServerError, errors.New("erro no banco de dados")
		}
	}

	js, err := json.Marshal(send)
	if err != nil {
		log.Println("SendNotification Marshal() %w", err)
		return http.StatusInternalServerError, errors.New("erro interno")
	}

	req, err := http.NewRequest(http.MethodPost, config.C.GetString("http.send_notification"), bytes.NewBuffer(js))
	if err != nil {
		log.Println("SendNotification NewRequest() %w", err)
		return http.StatusInternalServerError, errors.New("erro da requisição externa")
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		log.Println("SendNotification NewRequest() %w", err)
		return http.StatusInternalServerError, errors.New("erro da requisição externa")
	}

	return http.StatusOK, nil
}

func ConvertJsonToTransction(data []byte) (int, *Transaction, error) {
	ts := &Transaction{}
	err := json.Unmarshal(data, ts)
	if err != nil {
		log.Println("ConvertJsonToTransction Unmarshal() %w", err)
		return http.StatusInternalServerError, nil, errors.New("erro interno")
	}
	return http.StatusOK, ts, nil
}

func SetData() (t, t2 time.Time) {
	t = time.Now()
	t2 = t.AddDate(0, 0, 30)
	return
}

// Functions utilizadas para fazer o estorno
func GetTranferParams(id int, s *Service) (int, *Reversal, error) {

	var query strings.Builder

	rs := &Reversal{}

	fmt.Fprintf(&query, `SELECT value, payer, payee FROM mydb.transaction where id = %d;`, id)

	rows, err := s.Repository.Select(query.String())
	if err != nil {
		log.Println(fmt.Errorf("GetTranferParams Select(): %w", err))
		return http.StatusInternalServerError, nil, errors.New("erro no banco de dados")
	}

	for rows.Next() {
		err := rows.Scan(&rs.Value, &rs.Payee, &rs.Payer)
		if err != nil {
			log.Println(fmt.Errorf("GetTranferParams rows.Scan(): %w", err))
			return http.StatusInternalServerError, nil, errors.New("erro no banco de dados")
		}
	}

	return http.StatusOK, rs, nil
}

func (s *Service) CheckTransferEfetue(id int) (int, error) {

	var query strings.Builder
	var types int

	fmt.Fprintf(&query, `SELECT count(id) FROM mydb.reversal where transaction_id = %d;`, id)

	rows, err := s.Repository.Select(query.String())
	if err != nil {
		log.Println(fmt.Errorf("CheckTransferEfetue Select(): %w", err))
		return http.StatusInternalServerError, errors.New("erro no banco de dados")
	}

	for rows.Next() {
		err := rows.Scan(&types)
		if err != nil {
			log.Println(fmt.Errorf("CheckTransferEfetue rows.Scan(): %w", err))
			return http.StatusInternalServerError, errors.New("erro no banco de dados")
		}
		if types != 0 {
			return http.StatusForbidden, errors.New("estorno já realizado")
		}
	}

	return http.StatusOK, nil
}

func (s *Service) ExecuteReversal(transactionID, payer, payee int) (int, error) {

	var query strings.Builder

	fmt.Fprintf(&query, `INSERT INTO mydb.reversal (transaction_id, payer, payee) VALUES (%d, %d, %d);`, transactionID, payer, payee)

	_, err := s.Repository.Insert(query.String())
	if err != nil {
		log.Println(fmt.Errorf("ExecuteReversal Select(): %w", err))
		return http.StatusInternalServerError, errors.New("erro no banco de dados")
	}

	return http.StatusOK, nil
}

func (s *Service) Reversal(data []byte) (int, string, error) {

	status, rs, err := ConvertJsonToReversal(data)
	if err != nil {
		return status, "", err
	}

	status, err = s.CheckPassword(rs.UserID, rs.Password)
	if err != nil {
		return status, "", err
	}

	status, err = s.CheckTransferEfetue(rs.TransferID)
	if err != nil {
		return status, "", err
	}

	status, ts, err := GetTranferParams(rs.TransferID, s)
	if err != nil {
		return status, "", err
	}

	status, nextBalance, err := s.CheckSufficientBalance(ts.Payer, ts.Value)
	if err != nil {
		return status, "", err
	}

	status, err = s.CheckApiResponse()
	if err != nil {
		return status, "", err
	}

	status, _, err = s.ExecuteTransaction(ts.Value, nextBalance, ts.Payer, ts.Payee, 2, time.Now(), time.Now())
	if err != nil {
		return status, "", err
	}

	status, err = s.ExecuteReversal(rs.TransferID, ts.Payee, ts.Payer)
	if err != nil {
		return status, "", err
	}

	return status, `{"message":"estorno efetuado com sucesso"}`, nil
}

func ConvertJsonToReversal(data []byte) (int, *ReversalRequest, error) {
	ts := &ReversalRequest{}
	err := json.Unmarshal(data, ts)
	if err != nil {
		log.Println("ConvertJsonToReversal Unmarshal() %w", err)
		return http.StatusInternalServerError, nil, errors.New("erro interno")
	}
	return http.StatusOK, ts, nil
}
