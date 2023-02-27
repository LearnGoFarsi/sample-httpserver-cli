package server

import (
	"fmt"
	"net/http"
)

type HttpServer struct {
	server http.Server
	mux    *http.ServeMux
}

func NewHttpServer(host string, port int) *HttpServer {
	return &HttpServer{
		server: http.Server{
			Addr: fmt.Sprintf("%s:%d", host, port),
		},
		mux: http.NewServeMux(),
	}
}

func (s *HttpServer) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(pattern, handler)
}

func (s *HttpServer) Start() error {
	s.server.Handler = s.mux
	return s.server.ListenAndServe()
}
