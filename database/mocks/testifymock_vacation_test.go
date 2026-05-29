package mocks

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/szaborics/vacations-api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteVacationByID(t *testing.T) {
	mockRepo := new(Database)

	tests := []struct {
		name        string
		vacationID  string
		mockSetup   func()
		expected    int64
		expectedErr string
	}{
		{
			name:       "Success - Vacation Deleted",
			vacationID: "615f5f5f5f5f5f5f5f5f5f5f",
			mockSetup: func() {
				mockRepo.On("DeleteVacationByID", mock.Anything, "615f5f5f5f5f5f5f5f5f5f5f").Return(int64(1), nil)
			},
			expected:    int64(1),
			expectedErr: "",
		},
		{
			name:       "Success - No Vacation by that ID",
			vacationID: "515f5f5f5f5f5f5f5f5f5f5f",
			mockSetup: func() {
				mockRepo.On("DeleteVacationByID", mock.Anything, "515f5f5f5f5f5f5f5f5f5f5f").Return(int64(0), nil)
			},
			expected:    int64(0),
			expectedErr: "",
		},
		{
			name:       "Failure - Invalid Vacation ID",
			vacationID: "invalid-id",
			mockSetup: func() {
				mockRepo.On("DeleteVacationByID", mock.Anything, "invalid-id").Return(int64(0), errors.New("failed to initialize objectID: invalid object ID"))
			},
			expected:    int64(0),
			expectedErr: "failed to initialize objectID: invalid object ID",
		},
		{
			name:       "Failure - Database Error",
			vacationID: "315f5f5f5f5f5f5f5f5f5f5f",
			mockSetup: func() {
				mockRepo.On("DeleteVacationByID", mock.Anything, "315f5f5f5f5f5f5f5f5f5f5f").Return(int64(0), errors.New("database error"))
			},
			expected:    int64(0),
			expectedErr: "database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			deletedCount, err := mockRepo.DeleteVacationByID(context.Background(), tt.vacationID)

			assert.Equal(t, tt.expected, deletedCount)
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetVacationByID(t *testing.T) {
	mockRepo := new(Database)
	oid, err := primitive.ObjectIDFromHex("615f5f5f5f5f5f5f5f5f5f5f")
	if err != nil {
		t.Errorf("Failed to set object ID")
	}

	mockPOI := []models.POI{{
		Name:        "Cathedral di Santa Maria, Duomo",
		Description: "The Domed basilica of Santa Maria del Fiore",
	}}

	mockVacation := &models.VacationDAO{
		ID:               oid,
		City:             "Firenze",
		Country:          "Italy",
		Food:             "Bistecca a'la Fiurentino",
		Wine:             "Chianti",
		PointsOfInterest: mockPOI,
	}

	tests := []struct {
		name        string
		vacationID  string
		mockSetup   func()
		expected    *models.VacationDAO
		expectedErr error
	}{
		{
			name:       "Success - Get Vacation",
			vacationID: "615f5f5f5f5f5f5f5f5f5f5f",
			mockSetup: func() {
				mockRepo.On("GetVacationByID", mock.Anything, "615f5f5f5f5f5f5f5f5f5f5f").Return(mockVacation, nil)
			},
			expected:    mockVacation,
			expectedErr: nil,
		},
		{
			name:       "Failure - Vacation Not Found",
			vacationID: "515f5f5f5f5f5f5f5f5f5f5f",
			mockSetup: func() {
				mockRepo.On("GetVacationByID", mock.Anything, "515f5f5f5f5f5f5f5f5f5f5f").Return(nil, errors.New("vacation 515f5f5f5f5f5f5f5f5f5f5f not found"))
			},
			expected:    nil,
			expectedErr: errors.New("vacation 515f5f5f5f5f5f5f5f5f5f5f not found"),
		},
		{
			name:       "Failure - Invalid Object ID",
			vacationID: "415f5f5f5f5f5f5f5f5f5f5f",
			mockSetup: func() {
				mockRepo.On("GetVacationByID", mock.Anything, "415f5f5f5f5f5f5f5f5f5f5f").Return(nil, errors.New("failed to initialize objectID: invalid object ID"))
			},
			expected:    nil,
			expectedErr: errors.New("failed to initialize objectID: invalid object ID"),
		},
		{
			name:       "Failure - Database Failure",
			vacationID: "315f5f5f5f5f5f5f5f5f5f5f",
			mockSetup: func() {
				mockRepo.On("GetVacationByID", mock.Anything, "315f5f5f5f5f5f5f5f5f5f5f").Return(nil, errors.New("database failure"))
			},
			expected:    nil,
			expectedErr: errors.New("database failure"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			vacation, err := mockRepo.GetVacationByID(context.Background(), tt.vacationID)

			if tt.expectedErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, vacation)
			} else {
				assert.EqualError(t, err, tt.expectedErr.Error())
				assert.Nil(t, vacation)
			}
		})
	}
}
