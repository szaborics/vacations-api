package locations

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/szaborics/vacations-api/models"
)

func getLocationsByCountry(country string) []models.Location {

	var matchingPlaces []models.Location

	places := readFileLocations().Places

	for _, place := range places {
		if strings.EqualFold(place.Country, country) {

			matchingPlaces = append(matchingPlaces, place)

		}
	}
	return matchingPlaces

}

func readFileLocations() models.Locations {

	var filePath = "../data/locations.json"
	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)

	}

	var locations models.Locations
	err = json.Unmarshal(data, &locations)

	if err != nil {
		log.Fatalf("error parsing JSON %v", err)
	}

	return locations

}
