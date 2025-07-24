package usecase

import (
	"airline-voucher-seat-be/internal/models/domain"
	"airline-voucher-seat-be/internal/models/dto"
	"airline-voucher-seat-be/internal/models/view"
	"airline-voucher-seat-be/internal/repositories"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type VoucherUsecaseInterface interface {
	GenerateVoucher(ctxBackground context.Context, generateReq dto.VoucherGenerateReq) (*view.GeneratedVoucher, error)
	CheckVoucherIsExist(ctxBackground context.Context, checkReq dto.VoucherCheckReq) (*view.CheckVoucher, error)
	GetVouchers(ctxBackground context.Context) ([]domain.Voucher, error)
}

type VoucherUsecase struct {
	VoucherRepository  repositories.VoucherRepositoryInterface
	AircraftRepository repositories.AircraftRepositoryInterface
}

func NewVoucherUsecase(voucherRepository repositories.VoucherRepositoryInterface, aircraftRepository repositories.AircraftRepositoryInterface) VoucherUsecaseInterface {
	return &VoucherUsecase{VoucherRepository: voucherRepository, AircraftRepository: aircraftRepository}
}

func (u *VoucherUsecase) GenerateVoucher(ctx context.Context, generateReq dto.VoucherGenerateReq) (*view.GeneratedVoucher, error) {
	// Get Layout
	layout, err := u.AircraftRepository.GetSeatLayoutByAircraft(ctx, generateReq.Aircraft)
	if err != nil {
		return nil, errors.New("[VoucherUsecase] failed to get seat layout: " + err.Error())
	}

	// Generate seat list
	seatList := generateSeatList(layout)

	// Random seat generation logic can be added here
	seats, err := generateRandomSeats(seatList, 3)
	if err != nil {
		return nil, errors.New("[VoucherUsecase] failed to generate random seats: " + err.Error())
	}

	// Store the voucher in the database
	voucherDomain := domain.Voucher{
		CrewName:     generateReq.Name,
		CrewID:       generateReq.ID,
		FlightNumber: generateReq.FlightNumber,
		FlightDate:   generateReq.Date,
		AircraftType: generateReq.Aircraft,
		Seat1:        seats[0],
		Seat2:        seats[1],
		Seat3:        seats[2],
		CreatedAt:    time.Now(),
	}

	err = u.VoucherRepository.Create(ctx, voucherDomain)
	if err != nil {
		return nil, errors.New("[VoucherUsecase] failed to create voucher: " + err.Error())
	}

	return &view.GeneratedVoucher{Seats: seats}, nil
}

func generateSeatList(layout domain.SeatLayout) []string {
	seats := strings.Split(layout.SeatsPerRow, ",")
	var result []string
	for row := layout.RowStart; row <= layout.RowEnd; row++ {
		for _, seat := range seats {
			result = append(result, fmt.Sprintf("%d%s", row, seat))
		}
	}
	return result
}

func generateRandomSeats(seatList []string, count int) ([]string, error) {
	if len(seatList) < count {
		return nil, fmt.Errorf("[GenerateRandomSeats] requested %d seat(s), but only %d available", count, len(seatList))
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(seatList), func(i, j int) {
		seatList[i], seatList[j] = seatList[j], seatList[i]
	})

	return seatList[:count], nil
}

func (u *VoucherUsecase) CheckVoucherIsExist(ctx context.Context, checkReq dto.VoucherCheckReq) (*view.CheckVoucher, error) {
	exists, err := u.VoucherRepository.CheckVoucherIsExist(ctx, checkReq)
	if err != nil {
		return nil, errors.New("[VoucherUsecase] failed to check voucher: " + err.Error())
	}

	return &view.CheckVoucher{Exists: exists}, nil
}

func (u *VoucherUsecase) GetVouchers(ctx context.Context) ([]domain.Voucher, error) {
	vouchers, err := u.VoucherRepository.GetVouchers(ctx)
	if err != nil {
		return nil, errors.New("[VoucherUsecase] failed to get vouchers: " + err.Error())
	}

	return vouchers, nil
}
