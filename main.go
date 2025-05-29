// @title NebosukeRouteAPI
// @description ねぼすけあんないのAPI

package main

import (
	"log"
	"os"

	_ "github.com/Izana180/NebosukeTrainRouteSearch_backend/docs"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/handler"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/interface/repositoryimpl"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Println("env file not found")
	}

	port := os.Getenv("APP_PORT")

	stationRepo := repositoryimpl.NewStationRepository()
	stationUc := usecase.NewStationUsecase(stationRepo)
	stationHandler := handler.NewStationHandler(stationUc)

	routeRepo := repositoryimpl.NewRouteRepository()
	routeUc := usecase.NewRouteUsecase(routeRepo)
	routeHandler := handler.NewRouteHandler(routeUc)

	r := gin.Default()
	r.GET("/stations", stationHandler.GetAllStations)
	r.GET("/routesearch", routeHandler.GetRoute)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("server successfully runnning on " + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
