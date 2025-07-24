package controllers

import (
	"airline-voucher-seat-be/internal/models/dto"
	"airline-voucher-seat-be/internal/models/view"
	"airline-voucher-seat-be/internal/usecase"
	"airline-voucher-seat-be/middlewares"
	"context"

	"github.com/gofiber/fiber/v2"
)

type VoucherControllerInterface interface {
	GenerateVoucher(ctxBackground context.Context) func(*fiber.Ctx) error
	CheckVoucher(ctxBackground context.Context) func(*fiber.Ctx) error
	GetVouchers(ctxBackground context.Context) func(*fiber.Ctx) error
}

type VoucherController struct {
	VoucherUsecase usecase.VoucherUsecaseInterface
}

func NewVoucherController(voucherUsecase usecase.VoucherUsecaseInterface) VoucherControllerInterface {
	return &VoucherController{VoucherUsecase: voucherUsecase}
}

func (u *VoucherController) GenerateVoucher(ctxBackground context.Context) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var generateReq dto.VoucherGenerateReq
		if err := ctx.BodyParser(&generateReq); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
				Success: false,
				Message: "Invalid request body",
				Error:   err.Error(),
			})
		}
		if err := middlewares.ValidateBody(&generateReq); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
				Success: false,
				Message: "Invalid request body",
				Error:   err,
			})
		}

		seats, err := u.VoucherUsecase.GenerateVoucher(ctxBackground, generateReq)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
				Success: false,
				Message: "Failed to generate voucher",
				Error:   err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(view.Response{
			Success: true,
			Message: "Voucher generated successfully",
			Data:    &seats,
		})
	}
}

func (u *VoucherController) CheckVoucher(ctxBackground context.Context) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var checkReq dto.VoucherCheckReq
		if err := ctx.BodyParser(&checkReq); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
				Success: false,
				Message: "Invalid request body",
				Error:   err.Error(),
			})
		}
		if err := middlewares.ValidateBody(&checkReq); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(view.Response{
				Success: false,
				Message: "Invalid request body",
				Error:   err,
			})
		}

		exists, err := u.VoucherUsecase.CheckVoucherIsExist(ctxBackground, checkReq)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
				Success: false,
				Message: "Failed to check voucher",
				Error:   err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(view.Response{
			Success: true,
			Message: "Voucher checked successfully",
			Data:    &exists,
		})
	}
}

func (u *VoucherController) GetVouchers(ctxBackground context.Context) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		vouchers, err := u.VoucherUsecase.GetVouchers(ctxBackground)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(view.Response{
				Success: false,
				Message: "Failed to get vouchers",
				Error:   err.Error(),
			})
		}

		return ctx.Status(fiber.StatusOK).JSON(view.Response{
			Success: true,
			Message: "Vouchers fetched successfully",
			Data:    &vouchers,
		})
	}
}
