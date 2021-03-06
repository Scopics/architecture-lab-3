package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Scopics/architecture-lab-3/server/restaurant"
)

type APIServer struct {
	Port    int
	Handler restaurant.HttpHandlerFunc
	server  *http.Server
}

func (s *APIServer) Start() error {
	if s.Handler == nil {
		return fmt.Errorf("http handler is undefined")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not specified")
	}

	handler := new(http.ServeMux)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	handler.HandleFunc("/restaurant", s.Handler)
	return s.server.ListenAndServe()
}

func (s *APIServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
