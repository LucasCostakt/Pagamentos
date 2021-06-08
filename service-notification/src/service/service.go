package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"service-notification/src"
	"service-notification/src/config"
)

func NewService() src.ServiceList {
	return &Service{}
}

func (s *Service) SendNotifcation(p []byte) (int, string, error) {

	status, err := s.CheckBody(p)
	if err != nil {
		return status, "", err
	}

	status, err = s.RequestApiSendNotification()
	if err != nil {
		return status, "", err
	}

	return status, `{"message":"Notificação enviada com sucesso"}`, nil

}

func (s *Service) CheckBody(p []byte) (int, error) {

	return http.StatusOK, nil
}

func (s *Service) RequestApiSendNotification() (int, error) {
	client := http.Client{}
	data := &struct {
		Message string `json:"message"`
	}{}

	req, err := http.NewRequest(http.MethodGet, config.C.GetString("http.api_notify"), nil)
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

	if data.Message != "Success" {
		return http.StatusUnauthorized, errors.New("não autorizado")
	}

	return http.StatusOK, nil
}
