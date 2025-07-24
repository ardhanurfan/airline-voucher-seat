package usecase

import (
	"airline-voucher-seat-be/internal/repositories/mocks" // Assumes mocks are in this path
	"airline-voucher-seat-be/internal/usecase"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// TestAircraftUsecase_GetAircraftTypes tests the GetAircraftTypes method of AircraftUsecase.
func TestAircraftUsecase_GetAircraftTypes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockAircraftRepo := mocks.NewMockAircraftRepositoryInterface(ctrl)
	aircraftUsecase := usecase.NewAircraftUsecase(mockAircraftRepo)

	// Define test cases in a table-driven format.
	cases := []struct {
		name          string
		stub          func(mockRepo *mocks.MockAircraftRepositoryInterface) // Function to set up mock expectations.
		expectedTypes []string                                              // The expected result in a success case.
		wantErr       bool                                                  // Whether an error is expected.
		expectedErr   error                                                 // The specific error expected.
	}{
		{
			name: "Success - Aircraft types found",
			stub: func(mockRepo *mocks.MockAircraftRepositoryInterface) {
				expected := []string{"Boeing 737", "Airbus A320"}
				mockRepo.EXPECT().GetAircraftTypes(ctx).Return(expected, nil).Times(1)
			},
			expectedTypes: []string{"Boeing 737", "Airbus A320"},
			wantErr:       false,
			expectedErr:   nil,
		},
		{
			name: "Error - Repository returns an error",
			stub: func(mockRepo *mocks.MockAircraftRepositoryInterface) {
				mockRepo.EXPECT().GetAircraftTypes(ctx).Return(nil, errors.New("database error")).Times(1)
			},
			expectedTypes: nil,
			wantErr:       true,
			expectedErr:   errors.New("database error"),
		},
	}

	// Iterate over each test case.
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			tc.stub(mockAircraftRepo)

			actualTypes, err := aircraftUsecase.GetAircraftTypes(ctx)

			if tc.wantErr {
				// If an error was expected:
				require.Error(t, err)                 // Assert that an error was returned.
				require.Equal(t, tc.expectedErr, err) // Assert that the error is the expected one.
			} else {
				// If no error was expected:
				require.NoError(t, err)                         // Assert that no error was returned.
				require.Equal(t, tc.expectedTypes, actualTypes) // Assert that the returned data is correct.
			}
		})
	}
}
