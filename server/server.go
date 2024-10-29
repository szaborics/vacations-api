package server

import (
	"fmt"
	"net/http"
	"strings"
)

// VacationStore ...
type VacationStore interface {
	GETVacation(city string) string
}

// VacationsServer ...
type VacationsServer struct {
	store VacationStore
}

// ServerHTTP ...
func (v *VacationsServer) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	city := strings.TrimPrefix(r.URL.Path, "/vacations/")
	fmt.Fprint(w, v.store.GETVacation(city))
}

// GETVacation ...
func GETVacation(city string) string {
	if city == "Pula" {
		return `{
            "city": "Pula",
            "country": "Croatia",
            "food": "Konoba",
            "wine": "Malvasia",
            "pointsOfInterest":[{
                "name": "Pula Colosseum",
                "description": "Best preserved Roman colosseum from the venitian empire"
            }]
        }`
	}
	if city == "Montalcino" {
		return `{
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
	}
	return ""

}
