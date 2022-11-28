package main

import (
	"BE-S2-B41/database"
	"BE-S2-B41/pkg/mysql"
	"BE-S2-B41/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	// Init godotenv here ...
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	// inisialisasi route
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // add this code

	fmt.Println("server runnung localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
