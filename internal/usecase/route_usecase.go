package usecase

import "github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/model"

type RouteUsecase interface {
	GetRoute(from, to, datetime string, via []string, isArrivalTime bool) (*model.Route, error)
}
