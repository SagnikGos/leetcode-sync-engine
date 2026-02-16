package main

import (
	"leetcode-sync-engine/internal/database"
	"leetcode-sync-engine/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Init()

	r := gin.Default()

	r.POST("/api/submission", handlers.HandleSubmission)

	r.Run(":8080")
}
