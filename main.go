package main

import (
	"DCC/configs"
	"DCC/models/databases"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// configs.LoadEnvVariables()
	configs.ConnectToDB()
	
	db := configs.GetDB() 
	db.AutoMigrate(&databases.Customers{})

	log.Print("Server Running at: http://localhost:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Information page")
		fmt.Fprintln(w, "Login : http://localhost:8080/login")
		fmt.Fprintln(w, "Logout : http://localhost:8080/logout")
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Login page - First Endpoint")
		fmt.Fprint(w, "Welcome To The Login Page")
	})
	
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Logout page - Second Endpoint")
		fmt.Fprint(w, "Welcome To The Logout Page")
	})
	http.ListenAndServe(":8080", nil)
}