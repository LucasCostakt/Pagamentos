package main

import (
	"log"
	"service-users/src/config"
	"service-users/src/repository"
	"service-users/src/server"
	"service-users/src/service"
)

func main() {
	config.LoadConfig("development")
	log.Println("Create new httpServer")
	http := server.NewServer()
	conn, _ := repository.OpenConnection()

	serv := server.NewServerList(service.NewService(repository.NewRepository(conn)))

	http.NewRoutes(serv)
	http.StartAPI()
}
