package main

import (
	"log"
	"service-notification/src/config"
	"service-notification/src/server"
	"service-notification/src/service"
)

func main() {
	config.LoadConfig("development")
	log.Println("Create new httpServer")
	http := server.NewServer()

	serv := server.NewServerList(service.NewService())

	http.NewRoutes(serv)
	http.StartAPI()
}
