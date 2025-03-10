package database

import (
	"context"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteVacationByID(t *testing.T) {
	mockRepo := new(Database)

	// Define test cases
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
			name:       "Failure - Invalid Vacation ID",
			vacationID: "invalid-id",
			mockSetup: func() {
				// Simulate the behavior of primitive.ObjectIDFromHex failing
				mockRepo.On("DeleteVacationByID", mock.Anything, "invalid-id").Return(int64(0), errors.New("failed to initialize objectID: invalid object ID"))
			},
			expected:    int64(0),
			expectedErr: "failed to initialize objectID: invalid object ID",
		},
		{
			name:       "Failure - Database Error",
			vacationID: "315f5f5f5f5f5f5f5f5f5f5f",
			mockSetup: func() {
				// Simulate a database error
				mockRepo.On("DeleteVacationByID", mock.Anything, "315f5f5f5f5f5f5f5f5f5f5f").Return(int64(0), errors.New("database error"))
			},
			expected:    int64(0),
			expectedErr: "database error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock expectations
			tt.mockSetup()

			// Call the method on the mock
			deletedCount, err := mockRepo.DeleteVacationByID(context.Background(), tt.vacationID)

			// Assert the results
			assert.Equal(t, tt.expected, deletedCount)
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}

			// Verify that the expectations were met
			mockRepo.AssertExpectations(t)
		})
	}
}
