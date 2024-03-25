package di

import (
	"MyFirstAPIgo/internal/domain"
	"MyFirstAPIgo/internal/handlers"
	"MyFirstAPIgo/internal/handlers/middleware"
	"MyFirstAPIgo/internal/pkg"
	"MyFirstAPIgo/internal/repository/postgres"
	"MyFirstAPIgo/internal/usecase"
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"os"
)

type Test_bill struct {
	err error

	secretKey   string
	databaseURL string

	connection pkg.Connection

	router http.Handler

	createBills *usecase.CreateBillUseCase

	billsRepository domain.BillRepository
	userRepository  domain.UserRepository

	postBillsHandler    *handlers.POSTBillsHandler
	postRegisterHandler *handlers.POSTRegisterHandler
}

func (c *Test_bill) DatabaseURL() string {
	if c.databaseURL == "" {
		c.databaseURL = os.Getenv("DATABASE_URL")
	}

	return c.databaseURL
}

func (c *Test_bill) Pool(ctx context.Context) pkg.Connection {
	if c.connection == nil {
		postgresPool, err := pgxpool.New(ctx, c.DatabaseURL())
		if err != nil {
			panic(err)
		}

		if err := postgresPool.Ping(ctx); err != nil {
			panic(err)
		}

	}

	return c.connection
}

func (c *Test_bill) SecretKey() string {
	if c.secretKey == "" {
		c.secretKey = os.Getenv("JWT_SECRET_KEY")
	}

	return c.secretKey
}

func (c *Test_bill) SetUserRepository(userRepository domain.UserRepository) {
	c.userRepository = userRepository
}

func NewTest_bill() *Test_bill {
	return &Test_bill{}
}

func (c *Test_bill) PostBillsHandler(ctx context.Context) *handlers.POSTBillsHandler {
	if c.postBillsHandler == nil {
		c.postBillsHandler = handlers.NewPOSTBillsHandler(c.CreateBills(ctx))
	}

	return c.postBillsHandler
}

func (c *Test_bill) CreateBills(ctx context.Context) *usecase.CreateBillUseCase {
	if c.createBills == nil {
		c.createBills = usecase.CreateNewBillUseCase(c.BillsRepository(ctx))
	}

	return c.createBills
}

func (c *Test_bill) BillsRepository(ctx context.Context) domain.BillRepository {
	if c.billsRepository == nil {
		c.billsRepository = postgres.NewBillRepository(c.Pool(ctx))
	}

	return c.billsRepository
}

func (c *Test_bill) SetBillsRepository(billsRepository domain.BillRepository) {
	c.billsRepository = billsRepository
}

func (c *Test_bill) HTTPRouter(ctx context.Context) http.Handler {
	if c.router != nil {
		return c.router
	}

	router := mux.NewRouter()
	router.Use(
		middleware.PanicMiddleware,
		// middleware.AuthMiddleware,
	)

	router.Handle("/bills", c.PostBillsHandler(ctx)).Methods(http.MethodPost)

	c.router = router

	return c.router
}
