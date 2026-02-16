package main

import (
	"leetcode-sync-engine/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/submission", handlers.HandleSubmission)

	r.Run(":8080")
}
