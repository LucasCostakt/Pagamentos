package src

import "net/http"

type ServerList interface {
	TesteLogica(w http.ResponseWriter, r *http.Request)
}

type Storage interface {
	NewRoutes(serv ServerList)
	StartAPI()
}
