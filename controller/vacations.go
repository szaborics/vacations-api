package controller

//handles the HTTP resquest/responses, calls on the service layer to fulfill the request with the business logic applied,
// and surface any errors it receives from the service layer to the user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/szaborics/vacations-api/models"
	"github.com/szaborics/vacations-api/services"
)

// VacationController handles vacation requests
type VacationController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetFiltered(w http.ResponseWriter, r *http.Request)
	// Post(w http.ResponseWriter, r *http.Request)
	// Delete(w http.ResponseWriter, r *http.Request)
	HandleRoot(w http.ResponseWriter, r *http.Request)
}

// VacationControllerImpl implements vacation service
type VacationControllerImpl struct {
	vacationService services.VacationService
}

// NewVacationController returns address of vacation controller instance
func NewVacationController() VacationController {
	return &VacationControllerImpl{
		vacationService: services.NewVacationService(),
	}
}

var _ VacationController = (*VacationControllerImpl)(nil)

// GetAll Handles http request to get all vacations
func (v *VacationControllerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	var allVacations []models.VacationDTO

	allVacations, err := v.vacationService.GetAll(r.Context())
	if err != nil {
		sendJSONError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//If failed to encode
	if err := json.NewEncoder(w).Encode(allVacations); err != nil {
		sendJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
	return
}

// GetFiltered handles http requests to get a filtered list of vacations
func (v *VacationControllerImpl) GetFiltered(w http.ResponseWriter, r *http.Request) {

	var filter models.VacationFilter
	var filteredVacations []models.VacationDTO

	filter.Country = r.URL.Query().Get("country")
	filter.City = r.URL.Query().Get("city")

	filteredVacations, err := v.vacationService.Get(r.Context(), &filter)
	if err != nil {
		sendJSONError(w, http.StatusNotFound, "No Vacations found by that filter")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// if failed to encode
	if err := json.NewEncoder(w).Encode(filteredVacations); err != nil {
		sendJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
	return

}

// HandleRoot /
func (v *VacationControllerImpl) HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Vacations API!")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Welcome to my Vacations API!")
	return
}

func sendJSONError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
