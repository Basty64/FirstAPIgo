package entities

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	handler    http.Handler
}

//func NewServer() *Server {
//	return &Server{httpServer: httpServer, handler: handler}
//}

func (s *Server) Run(port string) error {

	err := http.ListenAndServe(port, s.handler)
	if err != nil {
		return err
	}
	return nil

}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
