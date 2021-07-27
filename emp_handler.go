package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	DbCrud.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}

func GetEmployees(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var emps []Employee
	DbCrud.Find(&emps)
	json.NewEncoder(w).Encode(emps)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var emps Employee
	DbCrud.First(&emps, mux.Vars(r)["ByID"])
	json.NewEncoder(w).Encode(emps)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var emps Employee
	DbCrud.First(&emps, mux.Vars(r)["ByID"])
	json.NewDecoder(r.Body).Decode(&emps)
	DbCrud.Save(&emps)
	json.NewEncoder(w).Encode(emps)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var emps Employee
	DbCrud.Delete(&emps, mux.Vars(r)["ByID"])
	json.NewEncoder(w).Encode("Employee Successfully Deleted")
}