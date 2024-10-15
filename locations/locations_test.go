package locations

import (
	"fmt"
	"testing"
	"reflect"
	"github.com/szaborics/vacations-api/models"
)

func TestGetLocationsByCountry(t *testing.T) {
	got := getLocationsByCountry("croatia")
	jsonData := models.Location {
		// Places: []models.Location {
		// {
            City: "Pula",
            Country: "Croatia",
            Food: "",
            Wine: "Malvasia",
            PointsOfInterest: models.POI{
                Name: "Pula Colosseum",
                Description: "Best preserved Roman colosseum from the venitian empire",
			},
        }
	// },
// }
	var want []models.Location 
	want = append(want, jsonData)
	
	if !reflect.DeepEqual(got, want) {
	    t.Errorf("got %v, want %v", got, want)
	}
	fmt.Println(got)

}




	// want:= {
	//         "city": "Pula",
	//         "country": "Croatia",
	//         "food": "",
	//         "wine": "Malvasia",
	//         "pointsOfInterest":{
	//             "name": "Pula Colosseum",
	//             "description": "Best preserved roan colosseam from the venitian empire"
	//         }
	//     }
	// if got!=want {
	// 	t.Errorf("Expected %s got %s", want, got )
	// }
func TestReadFileLocations(t *testing.T) {
	got := readFileLocations()

	fmt.Println(got)
}
