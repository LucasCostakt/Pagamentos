package server

import (
	"log"
	"net/http"
	"service-notification/src"
	"service-notification/src/config"
)

func NewServer() src.Storage {
	return new(httpServer)
}

//Init new routes
func (h *httpServer) NewRoutes(serv src.ServerList) {
	log.Println("Init Routes")
	router := http.NewServeMux()

	//Create the endpoins
	router.Handle("/", http.HandlerFunc(serv.SendNotification))

	h.Handler = router
}

//Run server
func (h *httpServer) StartAPI() {

	//load configs
	port := config.C.GetString("http.http_port")

	log.Println("Start API")
	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, h); err != nil {
		log.Fatal("init server error in StartApi(), ", err)
	}
}

func NewServerList(service src.ServiceList) src.ServerList {
	return &Server{
		Service: service,
	}

}
