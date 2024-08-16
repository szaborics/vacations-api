package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", basicHandler)
	router.Get("/vacations", getVacations)
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome to my vacation api!\n"))
}

func getVacations(w http.ResponseWriter, r *http.Request) {
	// read json values for vacations from file
	filePath := "./data/locations.json"

	jsonLocations, readErr := ioutil.ReadFile(filePath)
	if readErr != nil {
		fmt.Println(readErr)
	}

	w.Write([]byte(jsonLocations))

}
