package handlers

import (
	"MyFirstAPIgo/internal/pkg"
	"MyFirstAPIgo/internal/usecase"
	"encoding/json"
	"errors"
	"net/http"
)

type GETUsersHandler struct {
	useCase *usecase.FindUserUseCase
}

func (handler *GETUsersHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	name := query.Get("name")

	users, err := handler.useCase.Handle(request.Context(), usecase.FindUserQuery{Name: name})
	if errors.Is(err, pkg.ErrInfraction) {
		http.Error(writer, err.Error(), http.StatusUnprocessableEntity)

		return
	}
	if err != nil {
		// TODO
	}

	if err := json.NewEncoder(writer).Encode(users); err != nil {
		// TODO
	}
}