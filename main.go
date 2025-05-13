package main

import (
	"log"
	"os"

	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/handler"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/interface/repositoryimpl"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("env file not fount")
    }

    port := os.Getenv("APP_PORT")

    repo := repositoryimpl.NewStationRepository()
    uc := usecase.NewStationUsecase(repo)
    stationHandler := handler.NewStationHandler(uc)

    r := gin.Default()
    r.GET("/stations", stationHandler.GetAllStations)

    log.Println("server successfully runnning on " + port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}