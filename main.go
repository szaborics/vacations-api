package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/szaborics/vacations-api/models"
)

// func main() {

// 	router := chi.NewRouter()
// 	router.Use(middleware.Logger)
// 	router.Get("/hello", basicHandler)
// 	router.Get("/vacations", getVacations)
// 	server := &http.Server{
// 		Addr:    ":3000",
// 		Handler: router,
// 	}
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		fmt.Println("failed to listen to server", err)
// 	}
// }

// func basicHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello, welcome to my vacation api!\n"))
// }

func getVacations() {
	// read json values for vacations from file
	filePath := "data/locations.json"

	jsonFile, err := os.Open(filePath)

	checkForErrors(err)

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	checkForErrors(err)

	var locations models.Locations
	err = json.Unmarshal(byteValue, &locations)

	checkForErrors(err)

	for _, location := range locations.Places {
		fmt.Println(location)
	}

}

func checkForErrors(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	getVacations()
}
