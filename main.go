package main

import (
	"db/database"
	"db/database/validators"
	"fmt"
)

func main() {
	db := database.NewDatabase()

	// Users table
	nameCol := database.NewStringColumn("name").AddValidator(&validators.RequiredValidator{}).AddValidator(&validators.StringMaxLengthValidator{MaxLength: 20})
	ageCol := database.NewIntColumn("age").AddValidator(&validators.IntMinMaxValidator{Min: 16, Max: 60})
	db.CreateTable("users", []database.Column{nameCol, ageCol})

	fmt.Println("Table created")

	usersTable, err := db.GetTable("users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := usersTable.Insert(map[string]interface{}{"name": interface{}("suhail"), "age": interface{}(20)}); err != nil {
		fmt.Println(err.Error())

		return
	}

	usersTable.Insert(map[string]interface{}{"name": interface{}("manoja"), "age": interface{}(30)})

	fmt.Println(usersTable.FindBy(map[string]interface{}{"age": interface{}(30)}))
	fmt.Println(usersTable.FetchAll())
}
