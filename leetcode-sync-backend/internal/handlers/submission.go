package handlers

import (
	"net/http"

	"leetcode-sync-engine/internal/models"

	"github.com/gin-gonic/gin"
)

func HandleSubmission(c *gin.Context) {
	var submission models.Submission

	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid submission payload",
		})
		return
	}

	// Temporary logging
	c.JSON(http.StatusOK, gin.H{
		"status": "received",
		"slug":   submission.Slug,
	})
}
