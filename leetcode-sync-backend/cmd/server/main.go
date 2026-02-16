package main

import (
	"github.com/gin-gonic/gin"
	"leetcode-sync-engine/internal/database"
	"leetcode-sync-engine/internal/handlers"
)

func main() {
	database.Init()

	r := gin.Default()

	r.POST("/api/submission", handlers.HandleSubmission)

	r.Run(":8080")
}
