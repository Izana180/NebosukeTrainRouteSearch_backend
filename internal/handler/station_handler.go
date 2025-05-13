package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/usecase"
)


type StationHandler struct {
	usecase usecase.StationUsecase
}

func NewStationHandler(u usecase.StationUsecase) *StationHandler {
	return &StationHandler{
		usecase: u,
	}
}

func (h *StationHandler) GetAllStations(c *gin.Context) {
	stations, err := h.usecase.GetAllStations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get stations data"})
		return
	}

	c.JSON(http.StatusOK, stations)
}