package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/szaborics/vacations-api/models"
	svcmocks "github.com/szaborics/vacations-api/services/mocks"
)

// Seed data matching scripts/seedVacations.mongodb.js
var emptyPOI = []models.POI{{Name: "", Description: ""}}

var italyVacations = []models.VacationDTO{
	{City: "Varenna", Country: "Italy", Food: "fish", PointsOfInterest: emptyPOI},
	{City: "Riomaggiore", Country: "Italy", Food: "anchovies and pesto", PointsOfInterest: emptyPOI},
	{City: "Vernazza", Country: "Italy", Food: "anchovies, pesto, gelato, focaccia", PointsOfInterest: emptyPOI},
	{City: "Corniglia", Country: "Italy", Food: "anchovies, pesto, gelato, focaccia", PointsOfInterest: emptyPOI},
	{City: "Firenze", Country: "Italy", Food: "Bistecca a'la Fiurentino", PointsOfInterest: emptyPOI},
	{City: "Montepulciano", Country: "Italy", Wine: "Montepulciano", PointsOfInterest: emptyPOI},
	{City: "Montalcino", Country: "Italy", Wine: "Brunello di Montalcino", PointsOfInterest: emptyPOI},
}

var croatiaVacations = []models.VacationDTO{
	{
		City: "Pula", Country: "Croatia", Food: "Konoba", Wine: "Malvasia",
		PointsOfInterest: []models.POI{{
			Name:        "Pula Colosseum",
			Description: "Best preserved Roman colosseum from the venitian empire",
		}},
	},
}

var singleVacation = &models.VacationDTO{
	ID: "615f5f5f5f5f5f5f5f5f5f5f", City: "Corniglia", Country: "Italy", Food: "anchovies, pesto, gelato, focaccia", PointsOfInterest: emptyPOI,
}

func allVacations() []models.VacationDTO {
	result := make([]models.VacationDTO, 0, len(italyVacations)+len(croatiaVacations))
	result = append(result, italyVacations...)
	result = append(result, croatiaVacations...)
	return result
}

func TestGetById(t *testing.T) {
	tests := []struct {
		name           string
		vacationID     string
		mockReturn     *models.VacationDTO
		mockError      error
		expectedStatus int
		checkBody      func(t *testing.T, body string)
	}{
		{
			name:           "valid id returns vacation",
			vacationID:     "615f5f5f5f5f5f5f5f5f5f5f",
			mockReturn:     singleVacation,
			expectedStatus: http.StatusOK,
			checkBody: func(t *testing.T, body string) {
				var got models.VacationDTO
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Equal(t, "615f5f5f5f5f5f5f5f5f5f5f", got.ID)
				assert.Equal(t, "Corniglia", got.City)
			},
		},
		{
			name:           "invalid or missing id returns 404",
			vacationID:     "000000000000000000000000",
			mockError:      errors.New("vacation not found"),
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := svcmocks.NewVacationService(t)
			mockSvc.On("GetVacationByID", mock.Anything, tt.vacationID).Return(tt.mockReturn, tt.mockError)

			ctrl := &VacationControllerImpl{vacationService: mockSvc}

			req := httptest.NewRequest(http.MethodGet, "/vacations/"+tt.vacationID, nil)
			req.SetPathValue("id", tt.vacationID)
			rr := httptest.NewRecorder()

			ctrl.GetByID(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			if tt.checkBody != nil {
				tt.checkBody(t, rr.Body.String())
			}
		})
	}
}

func TestGetFiltered(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		mockFilter     *models.VacationFilter
		mockReturn     []models.VacationDTO
		mockError      error
		expectedStatus int
		checkBody      func(t *testing.T, body string)
	}{
		{
			name:           "no filter returns all 8 vacations",
			url:            "/vacations",
			mockFilter:     &models.VacationFilter{},
			mockReturn:     allVacations(),
			expectedStatus: http.StatusOK,
			checkBody: func(t *testing.T, body string) {
				var got []models.VacationDTO
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Len(t, got, 8)
			},
		},
		{
			name:           "filter by country=Italy returns 7 vacations",
			url:            "/vacations?country=Italy",
			mockFilter:     &models.VacationFilter{Country: "Italy"},
			mockReturn:     italyVacations,
			expectedStatus: http.StatusOK,
			checkBody: func(t *testing.T, body string) {
				var got []models.VacationDTO
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Len(t, got, 7)
				for _, v := range got {
					assert.Equal(t, "Italy", v.Country)
				}
			},
		},
		{
			name:           "filter by country=Croatia returns Pula with its POI",
			url:            "/vacations?country=Croatia",
			mockFilter:     &models.VacationFilter{Country: "Croatia"},
			mockReturn:     croatiaVacations,
			expectedStatus: http.StatusOK,
			checkBody: func(t *testing.T, body string) {
				var got []models.VacationDTO
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Len(t, got, 1)
				assert.Equal(t, "Pula", got[0].City)
				assert.Equal(t, "Pula Colosseum", got[0].PointsOfInterest[0].Name)
			},
		},
		{
			name:           "filter by city=Pula returns matching vacation",
			url:            "/vacations?city=Pula",
			mockFilter:     &models.VacationFilter{City: "Pula"},
			mockReturn:     croatiaVacations,
			expectedStatus: http.StatusOK,
			checkBody: func(t *testing.T, body string) {
				var got []models.VacationDTO
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Len(t, got, 1)
				assert.Equal(t, "Pula", got[0].City)
			},
		},
		{
			name:           "filter with no results returns 200 with empty array",
			url:            "/vacations?country=France",
			mockFilter:     &models.VacationFilter{Country: "France"},
			mockReturn:     []models.VacationDTO{},
			expectedStatus: http.StatusOK,
			checkBody: func(t *testing.T, body string) {
				var got []models.VacationDTO
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Empty(t, got)
			},
		},
		{
			name:           "service error returns 500",
			url:            "/vacations",
			mockFilter:     &models.VacationFilter{},
			mockError:      errors.New("database connection failed"),
			expectedStatus: http.StatusInternalServerError,
			checkBody: func(t *testing.T, body string) {
				var got map[string]string
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Equal(t, "Internal server error", got["error"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := svcmocks.NewVacationService(t)
			mockSvc.On("Get", mock.Anything, tt.mockFilter).Return(tt.mockReturn, tt.mockError)

			ctrl := &VacationControllerImpl{vacationService: mockSvc}

			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rr := httptest.NewRecorder()

			ctrl.GetFiltered(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			if tt.checkBody != nil {
				tt.checkBody(t, rr.Body.String())
			}
		})
	}
}

// add Create Vacation unit tests
func TestCreateVacation(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		mockReturn     string
		mockError      error
		expectedStatus int
		checkBody      func(t *testing.T, body string)
	}{
		{
			name:           "successful creation of vacation, 201 status, id returned",
			body:           `{"city":"Eger","country":"Hungary"}`,
			mockReturn:     "615f5f5f5f5f5f5f5f5f5f5f",
			expectedStatus: 201,
			checkBody: func(t *testing.T, body string) {
				var got models.VacationDTO
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Equal(t, "615f5f5f5f5f5f5f5f5f5f5f", got.ID)
				assert.Equal(t, "Eger", got.City)
			},
		},
		{
			name:           "failed to create because City is not set",
			body:           `{"country":"Hungary"}`,
			expectedStatus: 422,
			checkBody: func(t *testing.T, body string) {
				var got map[string]string
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Equal(t, "City and Country are required", got["error"])
			},
		},
		{
			name:           "failed to create because Country is not set",
			body:           `{"city":"Eger"}`,
			expectedStatus: 422,
			checkBody: func(t *testing.T, body string) {
				var got map[string]string
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Equal(t, "City and Country are required", got["error"])
			},
		},
		{
			name:           "failed with a 500 error",
			body:           `{"city":"Eger","country":"Hungary"}`,
			mockReturn:     "",
			mockError:      errors.New("Database Failed"),
			expectedStatus: 500,
			checkBody: func(t *testing.T, body string) {
				var got map[string]string
				assert.NoError(t, json.Unmarshal([]byte(body), &got))
				assert.Equal(t, "Internal Server Error", got["error"])
			},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSvc := svcmocks.NewVacationService(t)
			if tt.mockReturn != "" || tt.mockError != nil {
				mockSvc.On("CreateVacation", mock.Anything, mock.Anything).Return(tt.mockReturn, tt.mockError)
			}
			ctrl := &VacationControllerImpl{vacationService: mockSvc}
			req := httptest.NewRequest(http.MethodPost, "/vactations", strings.NewReader(tt.body))
			rr := httptest.NewRecorder()

			ctrl.CreateVacation(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			if tt.checkBody != nil {
				tt.checkBody(t, rr.Body.String())
			}
		})
	}
}

func TestHandleRoot(t *testing.T) {
	mockSvc := svcmocks.NewVacationService(t)
	ctrl := &VacationControllerImpl{vacationService: mockSvc}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	ctrl.HandleRoot(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Header().Get("Content-Type"), "application/json")

	var got map[string]string
	assert.NoError(t, json.Unmarshal([]byte(rr.Body.String()), &got))
	assert.Equal(t, "Welcome to my Vacations API!", got["message"])
}
