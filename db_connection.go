package main

import "fmt"
import "gorm.io/driver/mysql"
import "gorm.io/gorm"



var DbCrud *gorm.DB
var urlDSN = "somi:12345@tcp(localhost:3306)/crudapidb?parseTime=true"
var err error

func DataMigration()  {
	DbCrud, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Connection Failed :(")
	}
	DbCrud.AutoMigrate(&Employee{})
}
