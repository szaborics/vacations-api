package services

import (
	"context"
	"fmt"
	"testing"
)

var vacationService = NewVacationService()

func TestGetAll(t *testing.T) {

	t.Run("Tests getting all vacations from the repository", func(t *testing.T) {

		got, err := vacationService.GetAll(context.Background())

		if err != nil {
			t.Errorf("Failed to get all vacations %v", err.Error())
		}
		fmt.Printf("got %v", got)

	})
}
