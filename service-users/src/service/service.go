package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"service-users/src"
	"strings"
)

func NewService(repo src.RepositoryList) src.ServiceList {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) InsertNewUser(data []byte) (int, string, error) {

	status, user, err := ConvertJsonToUser(data)
	if err != nil {
		return status, "", err
	}

	status, err = s.InsertUser(user.Name, user.Cpf_cnpj, user.Email, user.Password, user.User_type)
	if err != nil {
		return status, "", err
	}

	return status, `{"message":"transação efetuada com sucesso"}`, nil
}

func (s *Service) InsertUser(name, cpf_cnpj, email, password string, userType int) (int, error) {

	var query strings.Builder

	fmt.Fprintf(&query, `INSERT INTO mydb.users (name, cpf_cnpj, email, password, type)
	SELECT * FROM (SELECT '%s', '%s', '%s', '%s', %d) AS tmp
	WHERE NOT EXISTS (SELECT cpf_cnpj, email FROM mydb.users  WHERE cpf_cnpj = '%s' or email = '%s') LIMIT 1;`, name, cpf_cnpj, email, password, userType, cpf_cnpj, email)

	res, err := s.Repository.Insert(query.String())
	if err != nil {
		log.Println(fmt.Errorf("InsertUser Insert(): %w", err))
		return http.StatusInternalServerError, errors.New("erro no banco de dados")
	}
	check, err := res.RowsAffected()
	if err != nil {
		log.Println(fmt.Errorf("InsertUser RowsAffected(): %w", err))
		return http.StatusInternalServerError, errors.New("erro no banco de dados")
	}

	if check == 0 {
		return http.StatusForbidden, errors.New("email ou cpf/cnpj já cadastrados")
	}

	return http.StatusOK, nil
}
func ConvertJsonToUser(data []byte) (int, *User, error) {
	user := &User{}

	err := json.Unmarshal(data, user)
	if err != nil {
		log.Println("ConvertJsonToReversal Unmarshal() %w", err)
		return http.StatusInternalServerError, nil, errors.New("erro interno")
	}

	return http.StatusOK, user, nil
}
