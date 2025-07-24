package controllers

import (
	"airline-voucher-seat-be/internal/models/view"
	"airline-voucher-seat-be/internal/usecase"
	"context"

	"github.com/gofiber/fiber/v2"
)

type AircraftControllerInterface interface {
	GetAircraftTypes(ctxBackground context.Context) func(*fiber.Ctx) error
}

type AircraftController struct {
	AircraftUsecase usecase.AircraftUsecaseInterface
}

func NewAircraftController(aircraftUsecase usecase.AircraftUsecaseInterface) AircraftControllerInterface {
	return &AircraftController{AircraftUsecase: aircraftUsecase}
}

func (u *AircraftController) GetAircraftTypes(ctxBackground context.Context) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		aircraftTypes, err := u.AircraftUsecase.GetAircraftTypes(ctxBackground)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
				Success: false,
				Message: "Failed to fetch aircraft types",
				Error:   err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(view.Response{
			Success: true,
			Message: "Aircraft types fetched successfully",
			Data:    aircraftTypes,
		})
	}
}
