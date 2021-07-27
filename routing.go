package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/employee/{ByID}", GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employee/{ByID}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{ByID}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080",r))
}


