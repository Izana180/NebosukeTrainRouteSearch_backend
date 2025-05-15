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

// GetAllStations godoc
// @Summary 駅一覧の取得
// @Description 外部APIから取得した全駅一覧を返す
// @Tags stations
// @Produce json
// @Success 200 {array} model.Station
// @Router /stations [get]
func (h *StationHandler) GetAllStations(c *gin.Context) {
	stations, err := h.usecase.GetAllStations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get stations data"})
		return
	}

	c.JSON(http.StatusOK, stations)
}