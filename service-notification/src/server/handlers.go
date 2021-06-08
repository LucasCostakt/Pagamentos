package server

import (
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) SendNotification(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("erro ao ler a requisição", err)
			http.Error(w, "erro ao ler a requisição", http.StatusInternalServerError)
			return
		}
		status, message, err := s.Service.SendNotifcation(body)
		if err != nil {
			http.Error(w, `{"message":`+`"`+err.Error()+`"`+`}`, status)
		}
		w.WriteHeader(status)
		w.Write([]byte(message))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
