package entities

import (
	"MyFirstAPIgo/internal/handlers"
	"MyFirstAPIgo/internal/repository/inmemory"
	"MyFirstAPIgo/internal/usecase"
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(ctx context.Context, addr string, postgresConnection string) (*Server, error) {
	router := mux.NewRouter()

	//pgxPool, err := pgxpool.New(ctx, postgresConnection)
	//if err != nil {
	//	return nil, err
	//}

	billRepository := inmemory.NewBillRepository()
	createBills := usecase.CreateNewBillUseCase(billRepository)

	router.Handle("/api/bill", handlers.NewPOSTBillsHandler(createBills))

	httpServer := &http.Server{Addr: addr, Handler: router}

	return &Server{httpServer: httpServer}, nil
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
