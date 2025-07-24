package usecase

import (
	"airline-voucher-seat-be/internal/models/domain"
	"airline-voucher-seat-be/internal/models/dto"
	"airline-voucher-seat-be/internal/repositories/mocks"
	"airline-voucher-seat-be/internal/usecase"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// TestVoucherUsecase_GenerateVoucher tests the GenerateVoucher method.
func TestVoucherUsecase_GenerateVoucher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockVoucherRepo := mocks.NewMockVoucherRepositoryInterface(ctrl)
	mockAircraftRepo := mocks.NewMockAircraftRepositoryInterface(ctrl)

	uc := usecase.NewVoucherUsecase(mockVoucherRepo, mockAircraftRepo)

	req := dto.VoucherGenerateReq{
		Name:         "John Doe",
		ID:           "12345",
		FlightNumber: "GA123",
		Date:         "2025-12-01",
		Aircraft:     "Boeing 737",
	}

	layout := domain.SeatLayout{
		RowStart:    1,
		RowEnd:      3,
		SeatsPerRow: "A,B,C",
	}

	cases := []struct {
		name    string
		stub    func()
		wantErr bool
	}{
		{
			name: "Success",
			stub: func() {
				mockAircraftRepo.EXPECT().GetSeatLayoutByAircraft(ctx, req.Aircraft).Return(layout, nil)
				mockVoucherRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Error - GetSeatLayoutByAircraft fails",
			stub: func() {
				mockAircraftRepo.EXPECT().GetSeatLayoutByAircraft(ctx, req.Aircraft).Return(domain.SeatLayout{}, errors.New("db error"))
			},
			wantErr: true,
		},
		{
			name: "Error - Create voucher fails",
			stub: func() {
				mockAircraftRepo.EXPECT().GetSeatLayoutByAircraft(ctx, req.Aircraft).Return(layout, nil)
				mockVoucherRepo.EXPECT().Create(ctx, gomock.Any()).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.stub()
			_, err := uc.GenerateVoucher(ctx, req)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestVoucherUsecase_CheckVoucherIsExist tests the CheckVoucherIsExist method.
func TestVoucherUsecase_CheckVoucherIsExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockVoucherRepo := mocks.NewMockVoucherRepositoryInterface(ctrl)
	uc := usecase.NewVoucherUsecase(mockVoucherRepo, nil)

	req := dto.VoucherCheckReq{
		FlightNumber: "GA123",
		Date:         "2025-12-01",
	}

	cases := []struct {
		name         string
		stub         func()
		expectedBool bool
		wantErr      bool
	}{
		{
			name: "Success - Voucher exists",
			stub: func() {
				mockVoucherRepo.EXPECT().CheckVoucherIsExist(ctx, req).Return(true, nil)
			},
			expectedBool: true,
			wantErr:      false,
		},
		{
			name: "Success - Voucher does not exist",
			stub: func() {
				mockVoucherRepo.EXPECT().CheckVoucherIsExist(ctx, req).Return(false, nil)
			},
			expectedBool: false,
			wantErr:      false,
		},
		{
			name: "Error - Repository fails",
			stub: func() {
				mockVoucherRepo.EXPECT().CheckVoucherIsExist(ctx, req).Return(false, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.stub()
			res, err := uc.CheckVoucherIsExist(ctx, req)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedBool, res.Exists)
			}
		})
	}
}

// TestVoucherUsecase_GetVouchers tests the GetVouchers method.
func TestVoucherUsecase_GetVouchers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockVoucherRepo := mocks.NewMockVoucherRepositoryInterface(ctrl)
	uc := usecase.NewVoucherUsecase(mockVoucherRepo, nil)

	vouchers := []domain.Voucher{
		{ID: 1, CrewName: "John Doe"},
		{ID: 2, CrewName: "Jane Doe"},
	}

	cases := []struct {
		name    string
		stub    func()
		wantErr bool
	}{
		{
			name: "Success",
			stub: func() {
				mockVoucherRepo.EXPECT().GetVouchers(ctx).Return(vouchers, nil)
			},
			wantErr: false,
		},
		{
			name: "Success - No vouchers found",
			stub: func() {
				mockVoucherRepo.EXPECT().GetVouchers(ctx).Return([]domain.Voucher{}, nil)
			},
			wantErr: false,
		},
		{
			name: "Error - Repository fails",
			stub: func() {
				mockVoucherRepo.EXPECT().GetVouchers(ctx).Return(nil, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.stub()
			_, err := uc.GetVouchers(ctx)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
