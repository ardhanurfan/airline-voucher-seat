package main

import (
	"airline-voucher-seat-be/config"
	"airline-voucher-seat-be/config/db"
	"airline-voucher-seat-be/internal/routes"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	ctx := context.Background()
	vouchersDB := db.InitVouchersDB(ctx)
	aircraftsDB := db.InitAircraftsDB(ctx)

	app := fiber.New()

	allowOrigin := config.GetEnv("CORS_ALLOW_ORIGINS", "http://localhost:3000")
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigin,
	}))

	api := app.Group("/api")
	routes.RestRoute(ctx, api, vouchersDB, aircraftsDB)

	app.Listen(":8080")
}
