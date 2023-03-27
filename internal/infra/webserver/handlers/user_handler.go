package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jeffersonbraster/apigo/internal/dto"
	"github.com/jeffersonbraster/apigo/internal/entity"
	"github.com/jeffersonbraster/apigo/internal/infra/database"
)


type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(UserDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: UserDB,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}