package vacations

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/szaborics/vacations-api/models"
)

func TestGetVacationsByCountry(t *testing.T) {
	//mock data
	poi1 := models.POI{
		Name:        "Pula Colosseum",
		Description: "Best preserved Roman colosseum from the venitian empire",
	}

	mockLocation := models.VacationDTO{
		City:             "Pula",
		Country:          "Croatia",
		Food:             "Konoba",
		Wine:             "Malvasia",
		PointsOfInterest: []models.POI{poi1},
	}

	mockData := models.VacationsDTO{
		Places: []models.VacationDTO{mockLocation},
	}

	got := getVacationsByCountry("croatia")
	want := mockData.Places

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got \n %v, want \n %v", got, want)
	}
	fmt.Println(got)

}

func TestReadFileVacations(t *testing.T) {
	got := readFileVacations()

	fmt.Println(got)
}
