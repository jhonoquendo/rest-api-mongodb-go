package user_route

import (
	userController "../../controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router  *mux.Router){
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")
}

