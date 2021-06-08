package server

import (
	"net/http"
	"service-notification/src"
)

type httpServer struct {
	http.Handler
}

type Server struct {
	Service src.ServiceList
}
