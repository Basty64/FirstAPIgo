package handlers

import (
	"MyFirstAPIgo/internal/pkg"
	"MyFirstAPIgo/internal/usecase"
	"encoding/json"
	"errors"
	"net/http"
)

type POSTRegisterHandler struct {
	useCase *usecase.CreateUserUseCase
}

func NewPOSTRegisterHandler(useCase *usecase.CreateUserUseCase) *POSTRegisterHandler {
	return &POSTRegisterHandler{useCase: useCase}
}

type POSTRegisterRequest struct {
	// Username is "user"
	Username string `json:"name"`
	// Password is "password"
	Password string `json:"password"`
}

func (handler *POSTRegisterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var r POSTRegisterRequest
	if err := json.NewDecoder(request.Body).Decode(&r); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)

		return
	}

	// Сохранить пользователя
	token, err := handler.useCase.Register(
		request.Context(),
		usecase.CreateUserCommand{
			Username: r.Username,
			Password: []byte(r.Password),
		},
	)
	var violation *pkg.Violation
	if errors.As(err, &violation) {
		http.Error(writer, violation.Message, http.StatusUnprocessableEntity)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(http.StatusNoContent)
	writer.Header().Set("Authorization", token)
}
