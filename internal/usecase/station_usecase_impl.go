package usecase

import (
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/model"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/repository"
)

type stationUsecase struct {
	stationRepo repository.StationRepository
}

func NewStationUsecase(repo repository.StationRepository) StationUsecase {
	return &stationUsecase{
		stationRepo: repo,
	}
}

func (u *stationUsecase) GetAllStations() ([]*model.Station, error) {
	return u.stationRepo.FetchAllStations()
}