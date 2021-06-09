package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"service-transaction/src"
	"service-transaction/src/config"
	"service-transaction/src/repository"
	"service-transaction/src/service"
	"testing"
)

var transferID int

func newServiceClientConnection() (src.ServiceList, *sql.DB) {
	config.LoadConfig("test")
	conn, err := repository.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	serv := service.NewService(repository.NewRepository(conn))
	return serv, conn
}

func TestCheckPassword(t *testing.T) {
	myStructTestResponse := []struct {
		name        string
		userID      int
		password    string
		want_status int
		want_error  string
	}{
		{
			name:        "Success Password",
			password:    "12345",
			userID:      1,
			want_status: http.StatusOK,
			want_error:  "",
		},
		{
			name:        "Password wrong",
			password:    "123455",
			userID:      1,
			want_status: http.StatusUnauthorized,
			want_error:  "senha ou usuário incorretos",
		},
		{
			name:        "User not exist",
			password:    "12345",
			userID:      156,
			want_status: http.StatusUnauthorized,
			want_error:  "senha ou usuário incorretos",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			res, conn := newServiceClientConnection()
			status, err := res.CheckPassword(tt.userID, tt.password)
			if err != nil {
				AssertResponsebody(t, err.Error(), tt.want_error)
			}
			AssertResponsebody(t, fmt.Sprint(status), fmt.Sprint(tt.want_status))

			conn.Close()
		})
	}
}

func TestCheckUserType(t *testing.T) {
	myStructTestResponse := []struct {
		name        string
		userID      int
		want_status int
		want_error  string
	}{
		{
			name:        "User have permission",
			userID:      2,
			want_status: http.StatusOK,
			want_error:  "",
		},
		{
			name:        "User not permission",
			userID:      1,
			want_status: http.StatusForbidden,
			want_error:  "permissão negada para realizar transferências",
		},
		{
			name:        "User not permission",
			userID:      145,
			want_status: http.StatusForbidden,
			want_error:  "usuário não encontrado",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			res, conn := newServiceClientConnection()
			status, err := res.CheckUserType(tt.userID)
			if err != nil {
				AssertResponsebody(t, err.Error(), tt.want_error)
			}
			AssertResponsebody(t, fmt.Sprint(status), fmt.Sprint(tt.want_status))

			conn.Close()
		})
	}
}

func TestCheckSufficientBalance(t *testing.T) {
	myStructTestResponse := []struct {
		name        string
		userID      int
		value       float64
		want_status int
		want_error  string
	}{
		{
			name:        "Sufficient balance",
			value:       10,
			userID:      3,
			want_status: http.StatusOK,
			want_error:  "",
		},
		{
			name:        "Insufficient balance",
			value:       10,
			userID:      2,
			want_status: http.StatusForbidden,
			want_error:  "saldo insuficiente",
		},
		{
			name:        "User not exist",
			value:       10,
			userID:      156,
			want_status: http.StatusForbidden,
			want_error:  "usuário não encontrado",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			res, conn := newServiceClientConnection()
			status, _, err := res.CheckSufficientBalance(tt.userID, tt.value)
			if err != nil {
				AssertResponsebody(t, err.Error(), tt.want_error)
			}
			AssertResponsebody(t, fmt.Sprint(status), fmt.Sprint(tt.want_status))

			conn.Close()
		})
	}
}

func TestCheckApiResponse(t *testing.T) {
	myStructTestResponse := []struct {
		name        string
		want_status int
		want_error  string
	}{
		{
			name:        "User have permission",
			want_status: http.StatusOK,
			want_error:  "",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			res, conn := newServiceClientConnection()
			status, err := res.CheckApiResponse()
			if err != nil {
				AssertResponsebody(t, err.Error(), tt.want_error)
			}
			AssertResponsebody(t, fmt.Sprint(status), fmt.Sprint(tt.want_status))

			conn.Close()
		})
	}
}

func TestExecuteTransaction(t *testing.T) {
	myStructTestResponse := []struct {
		name            string
		value           float64
		nextBalance     float64
		payer           int
		payee           int
		transactionType int
		want_status     int
		want_error      string
	}{
		{
			name:            "Users equals",
			value:           1,
			nextBalance:     4999,
			payer:           3,
			payee:           3,
			transactionType: 1,
			want_status:     http.StatusForbidden,
			want_error:      "usuários iguas",
		},
		{
			name:            "User not exist",
			value:           1,
			nextBalance:     4999,
			payer:           4566,
			payee:           100000,
			transactionType: 1,
			want_status:     http.StatusInternalServerError,
			want_error:      "erro no banco de dados",
		},
		{
			name:            "Sucess Transfer",
			value:           1,
			nextBalance:     4999,
			payer:           3,
			payee:           1,
			transactionType: 1,
			want_status:     http.StatusOK,
			want_error:      "",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			res, conn := newServiceClientConnection()

			transferTime, endDateReversal := service.SetData()
			status, tf, err := res.ExecuteTransaction(tt.value, tt.nextBalance, tt.payer, tt.payee, tt.transactionType, transferTime, endDateReversal)
			if err != nil {
				AssertResponsebody(t, err.Error(), tt.want_error)
			}
			AssertResponsebody(t, fmt.Sprint(status), fmt.Sprint(tt.want_status))
			transferID = tf
			conn.Close()
		})
	}
}

func TestSendNotification(t *testing.T) {
	myStructTestResponse := []struct {
		name        string
		transferID  int
		want_status int
		want_error  string
	}{
		{
			name:        "success send notification",
			transferID:  transferID,
			want_status: http.StatusOK,
			want_error:  "",
		},
		{
			name:        "notification not sent",
			transferID:  -1,
			want_status: http.StatusForbidden,
			want_error:  "transferencia não encontrada",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			res, conn := newServiceClientConnection()
			status, err := res.SendNotification(tt.transferID)
			if err != nil {
				AssertResponsebody(t, err.Error(), tt.want_error)
			}
			AssertResponsebody(t, fmt.Sprint(status), fmt.Sprint(tt.want_status))

			conn.Close()
		})
	}
}

func TestGetTranferParams(t *testing.T) {
	myStructTestResponse := []struct {
		name        string
		transferID  int
		want_status int
		want_error  string
	}{
		{
			name:        "without parameters",
			transferID:  0,
			want_status: http.StatusForbidden,
			want_error:  "transferencia não encontrado",
		},
		{
			name:        "transfer parameters ok",
			transferID:  transferID,
			want_status: http.StatusOK,
			want_error:  "",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			_, conn := newServiceClientConnection()

			s := &service.Service{Repository: repository.NewRepository(conn)}

			status, _, err := service.GetTranferParams(tt.transferID, s)
			if err != nil {
				AssertResponsebody(t, err.Error(), tt.want_error)
			}
			AssertResponsebody(t, fmt.Sprint(status), fmt.Sprint(tt.want_status))

			conn.Close()

		})
	}
}

func TestCheckTransferEfetue(t *testing.T) {
	myStructTestResponse := []struct {
		name        string
		transferID  int
		want_status int
		want_error  string
	}{
		{
			name:        "reversal can be made",
			transferID:  transferID,
			want_status: http.StatusOK,
			want_error:  "",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			res, conn := newServiceClientConnection()
			status, err := res.CheckTransferEfetue(tt.transferID)
			if err != nil {
				AssertResponsebody(t, err.Error(), tt.want_error)
			}
			AssertResponsebody(t, fmt.Sprint(status), fmt.Sprint(tt.want_status))
			conn.Close()

		})
	}
}

func TestExecuteReversal(t *testing.T) {
	myStructTestResponse := []struct {
		name        string
		payer       int
		payee       int
		transferID  int
		want_status int
		want_error  string
	}{
		{
			name:        "Users equals",
			payer:       3,
			payee:       3,
			transferID:  transferID,
			want_status: http.StatusForbidden,
			want_error:  "usuários iguas",
		},
		{
			name:        "Sucess Transfer",
			payer:       1,
			payee:       3,
			transferID:  transferID,
			want_status: http.StatusOK,
			want_error:  "",
		},
	}

	for _, tt := range myStructTestResponse {
		t.Run(tt.name, func(t *testing.T) {
			res, conn := newServiceClientConnection()

			status, err := res.ExecuteReversal(tt.transferID, tt.payer, tt.payee)
			if err != nil {
				AssertResponsebody(t, err.Error(), tt.want_error)
			}
			AssertResponsebody(t, fmt.Sprint(status), fmt.Sprint(tt.want_status))

			conn.Close()
		})
	}
}

func AssertResponsebody(t *testing.T, got, expectedResponse string) {
	t.Helper()
	if got != expectedResponse {
		t.Errorf("body is wrong, got %q want %q\n", got, expectedResponse)
	}
}
