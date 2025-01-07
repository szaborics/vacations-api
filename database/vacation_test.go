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

func validateResults(t testing.TB, got []models.VacationDAO, want []models.VacationDAO, err error) {
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

		validateResults(t, got, want, err)

	})
}

func TestInsertVacation(t *testing.T) {
	t.Run("test inserting a vacation to the vacations collection", func(t *testing.T) {

		id, err := primitive.ObjectIDFromHex("000000000000000000000000")
		if err != nil {
			t.Errorf("Failed to initialize ID")
		}

		mockPOI := []models.POI{{
			Name:        "test POI name",
			Description: "test POI Description",
		}}

		vacation := models.VacationDAO{
			ID:               id,
			City:             "Test city",
			Country:          "test country",
			Food:             "test food",
			Wine:             "test wine",
			PointsOfInterest: mockPOI,
		}

		ctx := context.Background()

		_, err = db.InsertVacation(ctx, vacation)

		if err != nil {
			t.Errorf("Failed to insert vacation: %v", err)
		}

	})

}
func TestDeleteVacationByID(t *testing.T) {

	t.Run("Tests deleting a single vacation document from the vacations collection", func(t *testing.T) {
		ctx := context.Background()
		id := "000000000000000000000000"

		want := int64(1)

		got, err := db.DeleteVacationByID(ctx, id)
		if err != nil {
			t.Fatalf("Failed to delete Vacation with id %s, error: %v", id, err)
		}

		if got != want {
			t.Errorf("The deleted count does not match expected, got %b, want %b", got, want)
		}

	})
}

func TestGetVacationByID(t *testing.T) {

	t.Run("Tests deleting a single vacation document from the vacations collection", func(t *testing.T) {
		ctx := context.Background()
		id := "6750a4e3454002bdba7e9bf8"
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			t.Fatalf("failed to set oid %v", err)
		}

		mockPOI := []models.POI{{
			Name:        "test POI name",
			Description: "test POI Description",
		}}

		vacation := &models.VacationDAO{
			ID:               oid,
			City:             "Test city",
			Country:          "test country",
			Food:             "test food",
			Wine:             "test wine",
			PointsOfInterest: mockPOI,
		}

		want := vacation
		got, err := db.GetVacationByID(ctx, id)
		if err != nil {
			t.Fatalf("Failed to get Vacation with id %s, error: %v", id, err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf(" got %v, wanted %v", got, want)
		}

	})
}
