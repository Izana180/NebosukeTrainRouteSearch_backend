// @title NebosukeRouteAPI
// @description ねぼすけあんないのAPI

package main

import (
	"log"
	"os"

    "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/handler"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/interface/repositoryimpl"
	"github.com/Izana180/NebosukeTrainRouteSearch_backend/internal/usecase"
    _ "github.com/Izana180/NebosukeTrainRouteSearch_backend/docs"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files"
)

func main() {
    if err := godotenv.Load("configs/.env"); err != nil {
        log.Println("env file not found")
    }

    port := os.Getenv("APP_PORT")

    repo := repositoryimpl.NewStationRepository()
    uc := usecase.NewStationUsecase(repo)
    stationHandler := handler.NewStationHandler(uc)

    r := gin.Default()
    r.GET("/stations", stationHandler.GetAllStations)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    log.Println("server successfully runnning on " + port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal(err)
    }
}