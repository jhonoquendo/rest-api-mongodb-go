package user_controller

import (
	m "../../models"
	userService "../../services/user.service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser m.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserta un usuario valido")
	}

	json.Unmarshal(reqBody, &newUser)
	userService.Create(newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := userService.Read()

	if err != nil{
		fmt.Fprintf(w, "Error al obtener usuarios")
	}

	json.NewEncoder(w).Encode(users)
}
