package repositoryimpl

import (
	"os"

	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/model"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/repository"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/infrastructure/odpt"
)
type stationRepository struct{
	apiKey string
}

func NewStationRepository() repository.StationRepository {
	return &stationRepository{
		apiKey: os.Getenv("ODPT_API_KEY"),
	}
}

func (r *stationRepository) FetchAllStations() ([]*model.Station, error) {
	rawStations, err := odpt.FetchStationsFromOdpt(r.apiKey)
	if err != nil {
		return nil, err
	}

	var stations []*model.Station
	for _, raw := range rawStations {
		stations = append(stations, &model.Station{
			ID: raw.ID,
			Name: raw.Title,
		})
	}

	return stations, nil
}