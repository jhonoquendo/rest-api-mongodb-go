package user_controller

import (
	m "../../models"
	userService "../../services/user.service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser m.User
	vars := mux.Vars(r)
	//userID, err := strconv.Atoi(vars["id"])
	userID := vars["id"]
	if userID == ""  {
		fmt.Fprintf(w, "Id vacio")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please ingrese dato valido")
		return
	}

	json.Unmarshal(reqBody, &updateUser)
	userService.Update(updateUser, userID)
	json.NewEncoder(w).Encode("Usuario actualizado con exito")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		fmt.Fprintf(w, "Id invalido")
		return
	}
	userService.Delete(userID)
	json.NewEncoder(w).Encode("Usuario eliminado con exito")
}
