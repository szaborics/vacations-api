package database

import (
	"context"
	"reflect"
	"testing"

	"github.com/szaborics/vacations-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetVacationCount(t *testing.T) {

	t.Run("test getting positive count by filter", func(t *testing.T) {

		ctx := context.Background()

		filter := models.VacationFilter{
			Country: "Italy",
		}

		got, err := db.GetVacationCount(ctx, &filter)
		want := int64(7)

		validateCountResults(t, got, want, err)

	})
	t.Run("test getting null results with filter", func(t *testing.T) {
		ctx := context.Background()

		filter := models.VacationFilter{
			Country: "Antarctica",
		}

		got, err := db.GetVacationCount(ctx, &filter)
		want := int64(0)

		validateCountResults(t, got, want, err)
	})

}

func validateCountResults(t testing.TB, got int64, want int64, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("results failed to be retrieved %v", err)
	}

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}

func validateResults(t testing.TB, got *[]models.VacationDAO, want *[]models.VacationDAO, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("results failed to be retrieved %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}

func TestGetFilteredVacations(t *testing.T) {

	t.Run("tests getting filtered vacations by country", func(t *testing.T) {

		ctx := context.Background()

		filter := models.VacationFilter{
			Country: "Croatia",
		}
		got, err := db.GetFilteredVacations(ctx, &filter)

		mockPOI := []models.POI{{
			Name:        "Pula Colosseum",
			Description: "Best preserved Roman colosseum from the venitian empire"},
		}

		specificID, err := primitive.ObjectIDFromHex("6719116d45173ca7957bfba5")

		if err != nil {
			t.Fatalf("failed to set specific ObjectId %v", err)
		}

		want := []models.VacationDAO{
			{
				ID:               specificID,
				City:             "Pula",
				Country:          "Croatia",
				Food:             "Konoba",
				Wine:             "Malvasia",
				PointsOfInterest: mockPOI,
			}}

		validateResults(t, &got, &want, err)

	})
}
