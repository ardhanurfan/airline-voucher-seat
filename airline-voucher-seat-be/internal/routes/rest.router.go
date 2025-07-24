package routes

import (
	"airline-voucher-seat-be/internal/controllers"
	"airline-voucher-seat-be/internal/repositories"
	"airline-voucher-seat-be/internal/usecase"
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func RestRoute(ctx context.Context, router fiber.Router, vouchersDB *sql.DB, aircraftsDB *sql.DB) {
	voucherRepository := repositories.NewVoucherRepository(vouchersDB)
	aircraftRepository := repositories.NewAircraftRepository(aircraftsDB)

	voucherUsecase := usecase.NewVoucherUsecase(voucherRepository, aircraftRepository)
	aircraftUsecase := usecase.NewAircraftUsecase(aircraftRepository)

	voucherController := controllers.NewVoucherController(voucherUsecase)

	aircraftController := controllers.NewAircraftController(aircraftUsecase)

	router.Post("/generate", voucherController.GenerateVoucher(ctx))
	router.Post("/check", voucherController.CheckVoucher(ctx))
	router.Get("/vouchers", voucherController.GetVouchers(ctx))
	router.Get("/aircrafts", aircraftController.GetAircraftTypes(ctx))
}
