package locations

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/szaborics/vacations-api/models"
)

func TestGetLocationsByCountry(t *testing.T) {
	//mock data
	poi1 := models.POI{
		Name:        "Pula Colosseum",
		Description: "Best preserved Roman colosseum from the venitian empire",
	}

	mockLocation := models.Location{
		City:             "Pula",
		Country:          "Croatia",
		Food:             "Konoba",
		Wine:             "Malvasia",
		PointsOfInterest: []models.POI{poi1},
	}

	mockData := models.Locations{
		Places: []models.Location{mockLocation},
	}

	got := getLocationsByCountry("croatia")
	want := mockData.Places

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got \n %v, want \n %v", got, want)
	}
	fmt.Println(got)

}

func TestReadFileLocations(t *testing.T) {
	got := readFileLocations()

	fmt.Println(got)
}
