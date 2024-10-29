package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubVacationStore struct {
	vacations map[string]string
}

func (v *StubVacationStore) GETVacation(city string) string {
	vacation := v.vacations[city]
	return vacation
}

func TestGETVacation(t *testing.T) {

	store := StubVacationStore{
		map[string]string{
			"Pula":`{
            "city": "Pula",
            "country": "Croatia",
            "food": "Konoba",
            "wine": "Malvasia",
            "pointsOfInterest":[{
                "name": "Pula Colosseum",
                "description": "Best preserved Roman colosseum from the venitian empire"
            }]
        }`,
		"Montalcino":`{
            "city": "Montalcino",
            "country": "Italy",
            "food": "",
            "wine": "Brunello di Montalcino",
            "pointsOfInterest":[
                {
                    "name": "",
                    "description": ""
                }
            ]
        }`,
		},
	}
	server := &VacationsServer{&store}
	t.Run("returns vacation Pula", func(t *testing.T) {

		request := newGetVacationByCityRequest("Pula")
		response := httptest.NewRecorder()

		server.ServerHTTP(response, request)

		got := response.Body.String()
		want := `{
            "city": "Pula",
            "country": "Croatia",
            "food": "Konoba",
            "wine": "Malvasia",
            "pointsOfInterest":[{
                "name": "Pula Colosseum",
                "description": "Best preserved Roman colosseum from the venitian empire"
            }]
        }`
		assertResponseBody(t, got, want)

	})
	t.Run("returns vacation Montalcino", func(t *testing.T) {
		request := newGetVacationByCityRequest("Montalcino")
		response := httptest.NewRecorder()

		server.ServerHTTP(response, request)

		got := response.Body.String()
		want := `{
            "city": "Montalcino",
            "country": "Italy",
            "food": "",
            "wine": "Brunello di Montalcino",
            "pointsOfInterest":[
                {
                    "name": "",
                    "description": ""
                }
            ]
        }`
		assertResponseBody(t, got, want)
	})
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func newGetVacationByCityRequest(city string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/vacations/%s", city), nil)
	return req

}
