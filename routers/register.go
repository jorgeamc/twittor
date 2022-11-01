package routers

import (
	"encoding/json"
	"github.com/jorgeamc/twittor/bd"
	"github.com/jorgeamc/twittor/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Bad request "+err.Error(), http.StatusBadRequest)
		return
	}
	if u.Email == "" {
		http.Error(w, "Email is not empty", http.StatusBadRequest)
	}

	if len(u.Password) < 6 {
		http.Error(w, "Password invalido longitud minima de 6", http.StatusBadRequest)
	}

	_, isUser, _ := bd.GetUser(u.Email)
	if isUser == true {
		http.Error(w, "Ya existe un usuario con el email registrado", http.StatusBadRequest)
		return
	}

	_, result, _ := bd.CreateRegister(u)
	if result == false {
		http.Error(w, "Error al crear el usuario", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
