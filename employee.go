package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Ename 			 string 	`json:"EmpName"`
	Edepartment      string  	`json:"EmpDep"`
	Emobile      	 float64 	`json:"EmpMob"`
	Eemail 			 string 	`json:"EmpEmail"`
	Esalary 		 float64 	`json:"EmpSalary"`

}
