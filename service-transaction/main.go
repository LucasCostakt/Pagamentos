package main

import (
	"log"
	"service-transaction/src/config"
	"service-transaction/src/repository"
	"service-transaction/src/server"
	"service-transaction/src/service"
)

func main() {
	config.LoadConfig("development")
	log.Println("Create new httpServer")
	http := server.NewServer()
	conn, err := repository.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	serv := server.NewServerList(service.NewService(repository.NewRepository(conn)))

	http.NewRoutes(serv)
	http.StartAPI()
}
