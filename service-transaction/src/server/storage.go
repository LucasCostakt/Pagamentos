package server

import (
	"net/http"
	"service-transaction/src"
)

type httpServer struct {
	http.Handler
}

type Server struct {
	Service src.ServiceList
}
