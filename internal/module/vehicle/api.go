package vehicle

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/qaultsabit/bengkel/domain"
	"github.com/qaultsabit/bengkel/internal/util"
)

type api struct {
	vehicleService domain.VehicleService
}

func NewApi(app *fiber.App, vehicleService domain.VehicleService) {
	api := api{
		vehicleService: vehicleService,
	}

	app.Get("v1/vehicle-histories", api.GetVehicleHistories)
}

func (a api) GetVehicleHistories(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30*time.Second)
	defer cancel()

	vin := ctx.Query("vin")
	if vin == "" {
		apiResponse := domain.ApiResponse{
			Code:    "400",
			Message: "INVALID PARAMETER",
		}
		util.ResponseInterceptor(c, &apiResponse)

		return ctx.Status(400).JSON(apiResponse)
	}

	apiResponse := a.vehicleService.FindHistorical(c, vin)
	util.ResponseInterceptor(c, &apiResponse)

	return ctx.Status(200).JSON(apiResponse)
}
