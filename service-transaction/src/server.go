package src

import "net/http"

type ServerList interface {
	Transaction(w http.ResponseWriter, r *http.Request)
	Reversal(w http.ResponseWriter, r *http.Request)
}

type Storage interface {
	NewRoutes(serv ServerList)
	StartAPI()
}
