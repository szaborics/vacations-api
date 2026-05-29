package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/szaborics/vacations-api/controller"
)

func main() {

	mux := http.NewServeMux()

	vacationController := controller.NewVacationController()

	mux.HandleFunc("GET /", vacationController.HandleRoot)
	mux.HandleFunc("GET /vacations", vacationController.GetFiltered)

	fmt.Println("Vacations API Server Listening on port 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
