package vehicle

import (
	"context"
	"time"

	"github.com/qaultsabit/bengkel/domain"
)

type service struct {
	vehicleRepository domain.VehicleRepository
	historyRepository domain.HistoryRepository
}

// FindHistorical implements domain.VehicleService.
func (s *service) FindHistorical(ctx context.Context, vin string) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, vin)
	if err != nil {
		return domain.ApiResponse{
			Code:    "500",
			Message: err.Error(),
		}
	}
	if vehicle == (domain.Vehicle{}) {
		return domain.ApiResponse{
			Code:    "404",
			Message: "VEHICLE NOT FOUND",
		}
	}
	histories, err := s.historyRepository.FindByVehicle(ctx, vehicle.ID)
	if err != nil {
		return domain.ApiResponse{
			Code:    "400",
			Message: err.Error(),
		}
	}

	var historiesData []domain.HistoryData
	for _, v := range histories {
		historiesData = append(historiesData, domain.HistoryData{
			VehicleID:   v.VehicleID,
			CustomerID:  v.CustomerID,
			PIC:         v.PIC,
			PlateNumber: v.PlateNumber,
			Notes:       v.Notes,
			ComeAt:      v.CreatedAt.Format(time.RFC822Z),
		})
	}
	result := domain.VehicleHistorical{
		ID:        vehicle.ID,
		VIN:       vehicle.VIN,
		Brand:     vehicle.Brand,
		Histories: historiesData,
	}

	return domain.ApiResponse{
		Code:    "200",
		Message: "APPROVED",
		Data:    result,
	}
}

func NewService(vehicleRepository domain.VehicleRepository,
	historyRepository domain.HistoryRepository) domain.VehicleService {
	return &service{vehicleRepository: vehicleRepository, historyRepository: historyRepository}
}
