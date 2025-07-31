package main

import (
	"go-training-application/database"
	"go-training-application/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.InitDB()

	r := gin.Default()
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/tasks", handlers.GetTasks)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.Run()
}
