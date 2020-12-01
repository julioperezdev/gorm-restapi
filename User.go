package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name  string
	Email string
}

func InitialMigration() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could connect to the database")
	}
	defer db.Close()
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could connect to the database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})

	fmt.Fprintf(w, "New User Successfully created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could connect to the database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, "User was deleted successfully")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could connect to the database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]
	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "User was updated successfully")
}
