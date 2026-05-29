package services

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	dbmocks "github.com/szaborics/vacations-api/database/mocks"
	"github.com/szaborics/vacations-api/models"
)

// Seed data matching scripts/seedVacations.mongodb.js
var emptyPOI = []models.POI{{Name: "", Description: ""}}

var italyDAOs = []models.VacationDAO{
	{City: "Varenna", Country: "Italy", Food: "fish", PointsOfInterest: emptyPOI},
	{City: "Riomaggiore", Country: "Italy", Food: "anchovies and pesto", PointsOfInterest: emptyPOI},
	{City: "Vernazza", Country: "Italy", Food: "anchovies, pesto, gelato, focaccia", PointsOfInterest: emptyPOI},
	{City: "Corniglia", Country: "Italy", Food: "anchovies, pesto, gelato, focaccia", PointsOfInterest: emptyPOI},
	{City: "Firenze", Country: "Italy", Food: "Bistecca a'la Fiurentino", PointsOfInterest: emptyPOI},
	{City: "Montepulciano", Country: "Italy", Wine: "Montepulciano", PointsOfInterest: emptyPOI},
	{City: "Montalcino", Country: "Italy", Wine: "Brunello di Montalcino", PointsOfInterest: emptyPOI},
}

var croatiaDAOs = []models.VacationDAO{
	{
		City: "Pula", Country: "Croatia", Food: "Konoba", Wine: "Malvasia",
		PointsOfInterest: []models.POI{{
			Name:        "Pula Colosseum",
			Description: "Best preserved Roman colosseum from the venitian empire",
		}},
	},
}

func TestGet(t *testing.T) {
	tests := []struct {
		name      string
		filter    *models.VacationFilter
		dbReturn  []models.VacationDAO
		dbError   error
		wantCount int
		wantErr   bool
	}{
		{
			name:      "empty filter returns all 8 vacations",
			filter:    &models.VacationFilter{},
			dbReturn:  append(append([]models.VacationDAO{}, italyDAOs...), croatiaDAOs...),
			wantCount: 8,
		},
		{
			name:      "filter by country=Italy returns 7 vacations",
			filter:    &models.VacationFilter{Country: "Italy"},
			dbReturn:  italyDAOs,
			wantCount: 7,
		},
		{
			name:      "filter by country=Croatia returns Pula",
			filter:    &models.VacationFilter{Country: "Croatia"},
			dbReturn:  croatiaDAOs,
			wantCount: 1,
		},
		{
			name:      "filter with no matching results returns empty slice",
			filter:    &models.VacationFilter{Country: "France"},
			dbReturn:  []models.VacationDAO{},
			wantCount: 0,
		},
		{
			name:    "database error is propagated",
			filter:  &models.VacationFilter{Country: "Italy"},
			dbError: errors.New("database connection failed"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := dbmocks.NewDatabase(t)
			mockDB.On("GetFilteredVacations", mock.Anything, tt.filter).Return(tt.dbReturn, tt.dbError)

			svc := &VacationServiceImpl{database: mockDB}
			result, err := svc.Get(context.Background(), tt.filter)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
				return
			}

			assert.NoError(t, err)
			assert.Len(t, result, tt.wantCount)

			// Verify DAO→DTO conversion: ID is stripped, named fields are preserved
			if tt.wantCount > 0 {
				assert.Equal(t, tt.dbReturn[0].City, result[0].City)
				assert.Equal(t, tt.dbReturn[0].Country, result[0].Country)
				assert.Equal(t, tt.dbReturn[0].Food, result[0].Food)
				assert.Equal(t, tt.dbReturn[0].Wine, result[0].Wine)
				assert.Equal(t, tt.dbReturn[0].PointsOfInterest, result[0].PointsOfInterest)
			}
		})
	}
}
