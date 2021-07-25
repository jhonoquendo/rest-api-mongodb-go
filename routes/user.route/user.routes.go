package user_route

import (
	userController "../../controllers/user.controller"
	"github.com/gorilla/mux"
)

func UserRoutes(router  *mux.Router){
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
}

