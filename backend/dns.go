package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //database variable
var err error

// Online db connection. Must be restarted. AWS was removed as no longer free.
//const DNS = "sql9624027:WHUUJxfxAJ@tcp(sql9.freemysqlhosting.net:3306)/sql9624027?charset=utf8&parseTime=True&loc=Local"

// local db connection -(fill in password and remove "<>", create "godb" schema in your database)
const DNS = "root:cici1998@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local"

// Initializes tables if necessary
func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("There was an error connecting", err.Error())
		panic("Cannot connect to DB")
	} else {
		fmt.Println("Connection to database successful.")
	}
	//Structs here get created as tables in db
	DB.AutoMigrate(&Stock{}, &UserStocks{}, &Credentials{}, &PortfolioHistory{})
}
