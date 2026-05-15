package main

import (
	"saints-api/controller"
	"saints-api/db"
	"saints-api/repository"
	"saints-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Repository Layer
	SaintRepository := repository.NewSaintRepository(dbConnection)
	//Usecase Layer
	SaintUseCase := usecase.NewSaintUseCase(SaintRepository)
	//Controllers Layer
	SaintController := controller.NewSaintController(SaintUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/saints", SaintController.GetSaints)
	server.GET("/saints/random", SaintController.GetRandomSaint)
	server.POST("/saint", SaintController.CreateSaint)
	server.GET("/saint/:saintId", SaintController.GetSaintById)

	server.Run(":8000")
}
