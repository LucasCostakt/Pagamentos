package server

import (
	"net/http"
	"service-users/src"
)

type httpServer struct {
	http.Handler
}

type Server struct {
	Service src.ServiceList
}
