package usecase

import (
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/model"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/domain/repository"
)

type routeUsecase struct {
	routeRepo repository.RouteRepository
}

func NewRouteUsecase(repo repository.RouteRepository) RouteUsecase {
	return &routeUsecase{
		routeRepo: repo,
	}
}

func (u *routeUsecase) GetRoute(from, to, datetime string, via []string, isArrivalTime bool) (*model.Route, error) {
	return u.routeRepo.FetchRouteWithNodeid(from, to, datetime, via, isArrivalTime)
}
