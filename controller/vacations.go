package controller

// handles the HTTP request/responses, calls on the service layer to fulfill the request with the business logic applied,
// and surfaces any errors it receives from the service layer to the caller

import (
	"encoding/json"
	"net/http"

	"github.com/szaborics/vacations-api/models"
	"github.com/szaborics/vacations-api/services"
)

// VacationController handles vacation requests
type VacationController interface {
	GetFiltered(w http.ResponseWriter, r *http.Request)
	HandleRoot(w http.ResponseWriter, r *http.Request)
}

// VacationControllerImpl implements VacationController
type VacationControllerImpl struct {
	vacationService services.VacationService
}

// NewVacationController returns an instance of VacationController wired with a live database service
func NewVacationController() VacationController {
	return &VacationControllerImpl{
		vacationService: services.NewVacationService(),
	}
}

var _ VacationController = (*VacationControllerImpl)(nil)

// GetFiltered handles GET /vacations — returns all vacations, optionally filtered by ?country= and/or ?city=
// Returns 200 with an empty array when no results match the filter.
func (v *VacationControllerImpl) GetFiltered(w http.ResponseWriter, r *http.Request) {

	filter := models.VacationFilter{
		Country: r.URL.Query().Get("country"),
		City:    r.URL.Query().Get("city"),
	}

	filteredVacations, err := v.vacationService.Get(r.Context(), &filter)
	if err != nil {
		sendJSONError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(filteredVacations); err != nil {
		sendJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}

}

// HandleRoot handles GET /
func (v *VacationControllerImpl) HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to my Vacations API!"})
}

func sendJSONError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
