package usecase

import (
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/model"
)

type StationUsecase interface {
	GetAllStations() ([]*model.Station, error)
}
