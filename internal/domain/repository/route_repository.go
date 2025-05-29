package repository

import "github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/model"

type RouteRepository interface {
	FetchRouteWithNodeid(from, to, datetime string, via []string, isArrivalTime bool) (*model.Route, error)
}
