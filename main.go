package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/qaultsabit/bengkel/internal/component"
	"github.com/qaultsabit/bengkel/internal/config"
	"github.com/qaultsabit/bengkel/internal/module/customer"
	"github.com/qaultsabit/bengkel/internal/module/history"
	"github.com/qaultsabit/bengkel/internal/module/vehicle"
)

func main() {
	conf := config.Get()
	dbConnection := component.GetDatabaseConnection(conf)

	customerRepository := customer.NewRepository(dbConnection)
	vehicleRepository := vehicle.NewRepository(dbConnection)
	historyRepository := history.NewRepository(dbConnection)

	customerService := customer.NewService(customerRepository)
	vehicleService := vehicle.NewService(vehicleRepository, historyRepository)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${locals:requestid}] ${ip} - ${method} ${status} ${path}\n",
	}))
	customer.NewApi(app, customerService)
	vehicle.NewApi(app, vehicleService)

	_ = app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
