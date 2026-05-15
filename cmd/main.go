package main

import (
	"saints-api/controller"
	"saints-api/db"
	"saints-api/middleware"
	"saints-api/repository"
	"saints-api/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

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
	
	// Protected POST route
	authorized := server.Group("/")
	authorized.Use(middleware.APIKeyAuth())
	{
		authorized.POST("/saint", SaintController.CreateSaint)
	}
	
	server.GET("/saint/:saintId", SaintController.GetSaintById)

	server.Run(":8000")
}
