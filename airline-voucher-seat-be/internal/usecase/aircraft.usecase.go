package usecase

import (
	"airline-voucher-seat-be/internal/repositories"
	"context"
)

type AircraftUsecaseInterface interface {
	GetAircraftTypes(ctxBackground context.Context) ([]string, error)
}

type AircraftUsecase struct {
	AircraftRepository repositories.AircraftRepositoryInterface
}

func NewAircraftUsecase(aircraftRepository repositories.AircraftRepositoryInterface) AircraftUsecaseInterface {
	return &AircraftUsecase{AircraftRepository: aircraftRepository}
}

func (u *AircraftUsecase) GetAircraftTypes(ctxBackground context.Context) ([]string, error) {
	return u.AircraftRepository.GetAircraftTypes(ctxBackground)
}
