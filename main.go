package main

import (
	"log"
	"net/http"
	"student_management/controllers"
	"student_management/models"

	"github.com/gorilla/mux"
)

func main() {

	models.InitDB("root:@tcp(127.0.0.1:3306)/student_management")

	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Index).Methods("GET")
	r.HandleFunc("/create", controllers.ShowCreateForm).Methods("GET")
	r.HandleFunc("/create", controllers.CreateStudent).Methods("POST")
	r.HandleFunc("/edit", controllers.ShowEditForm).Methods("GET")
	r.HandleFunc("/edit", controllers.UpdateStudent).Methods("POST")
	r.HandleFunc("/delete", controllers.DeleteStudent).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
