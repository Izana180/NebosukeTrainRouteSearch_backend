package handler

import (
	"net/http"

	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	usecase usecase.RouteUsecase
}

func NewRouteHandler(u usecase.RouteUsecase) *RouteHandler {
	return &RouteHandler{
		usecase: u,
	}
}

// GetRoute godoc
// @Summary 経路検索
// @Description 出発駅、到着駅、経由駅、出発時刻もしくは到着時刻を受け取り、経路を検索して返す
// @Tags routes
// @Produce json
// @Param from query string true "出発駅のnodeID(例: 00004464)"
// @Param to query string true "到着駅のnodeID(例: 00004254)"
// @Param datetime query string true "日付と時刻（例: 2025-05-28T10:00）"
// @Param isArrivalTime query bool false "datetimeが到着時刻かどうか(デフォルト: false)"
// @Param via query []string false "経由地(nodeID)※複数指定可能"
// @Success 200 {object} model.Route
// @Router /routesearch [get]
func (h *RouteHandler) GetRoute(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	datetime := c.Query("datetime")
	isArrivalTime := c.DefaultQuery("isArrivalTime", "false")
	// boolにキャスト
	isArrival := (isArrivalTime == "true")
	via := c.QueryArray("via")

	if from == "" || to == "" || datetime == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "params are not enough"})
		return
	}

	route, err := h.usecase.GetRoute(from, to, datetime, via, isArrival)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch route", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, route)
}
