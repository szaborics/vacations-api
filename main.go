package main

import (
	"fmt"
	"net/http"

	controller "github.com/szaborics/vacations-api/controller"
)

func main() {

	mux := http.NewServeMux()

	vacationController := controller.NewVacationController()

	mux.HandleFunc("GET /", vacationController.HandleRoot)
	// mux.HandleFunc("GET /vacations", vacationController.GetAll)
	mux.HandleFunc("GET /vacations", vacationController.GetFiltered)

	fmt.Println("Vacations API Server Listening on port 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server failed to start:", err)

	}

}
