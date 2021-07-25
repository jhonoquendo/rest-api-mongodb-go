package main

import (
	userR "./routes/user.route"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a mi API de Jhon Oquendo")
}

func main() {
	fmt.Println("conexion a Mongo DB")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	userR.UserRoutes(router)
	log.Fatal(http.ListenAndServe(":3000", router))
}
