package vacations

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/szaborics/vacations-api/models"
)

func getVacationsByCountry(country string) []models.VacationDTO {

	var filtered []models.VacationDTO

	places := readFileVacations().Places

	for _, place := range places {
		if strings.EqualFold(place.Country, country) {

			fmt.Println(place.String())
			filtered = append(filtered, place)

		}
	}
	return filtered

}

func readFileVacations() models.VacationsDTO {

	var filePath = "../../data/vacations.json"
	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)

	}

	var vacations models.VacationsDTO
	err = json.Unmarshal(data, &vacations)

	if err != nil {
		log.Fatalf("error parsing JSON %v", err)
	}

	return vacations

}
