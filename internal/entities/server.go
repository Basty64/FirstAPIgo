package entities

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {

	error := http.ListenAndServe(port, nil)
	if error != nil {
		return error
	}
	return nil

}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
