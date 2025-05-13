package repository

import "github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/model"

type StationRepository interface{
	FetchAllStations() ([]*model.Station, error)
}