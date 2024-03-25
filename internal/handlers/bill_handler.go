package handlers

import (
	"MyFirstAPIgo/internal/handlers/middleware"
	"MyFirstAPIgo/internal/pkg"
	"MyFirstAPIgo/internal/usecase"
	"encoding/json"
	"errors"
	"github.com/gofrs/uuid"
	"net/http"
)

type POSTBillsHandler struct {
	useCase *usecase.CreateBillUseCase
}

func NewPOSTBillsHandler(useCase *usecase.CreateBillUseCase) *POSTBillsHandler {
	return &POSTBillsHandler{useCase: useCase}
}

type POSTBillRequest struct {
	Name string `json:"name"`
}

type POSTBillResponse struct {
	id       uuid.UUID
	name     string
	isOpened bool
	userID   uuid.UUID
}

func (response *POSTBillResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		UserID   uuid.UUID `json:"user"`
		IsOpened bool      `json:"isOpened"`
	}{
		ID:       response.id,
		Name:     response.name,
		IsOpened: response.isOpened,
		UserID:   response.userID,
	})
}

func (handler *POSTBillsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	userID := middleware.UserIDFromContext(request.Context())

	var body POSTBillRequest
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	command := usecase.CreateCommand{
		Name:   body.Name,
		UserID: userID,
	}

	bill, err := handler.useCase.Handle(request.Context(), command)
	if errors.Is(err, pkg.ErrEmptyField) {
		http.Error(writer, err.Error(), http.StatusUnprocessableEntity)

		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	response := &POSTBillResponse{
		id:       bill.ID(),

		isOpened: !bill.IsClosed(),
		userID:   bill.UserID,
	}

	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}
}
