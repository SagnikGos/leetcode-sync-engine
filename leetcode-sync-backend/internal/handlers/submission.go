package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"leetcode-sync-engine/internal/database"
	"leetcode-sync-engine/internal/models"
	"leetcode-sync-engine/internal/utils"
)

func HandleSubmission(c *gin.Context) {
	var submission models.Submission

	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid submission payload",
		})
		return
	}

	codeHash := utils.HashCode(submission.Code)

	// Check duplicate submission
	var exists int
	err := database.DB.QueryRow(
		"SELECT COUNT(*) FROM submissions WHERE slug = ? AND code_hash = ?",
		submission.Slug,
		codeHash,
	).Scan(&exists)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
		return
	}

	if exists > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "duplicate",
			"slug":   submission.Slug,
		})
		return
	}

	// Insert submission
	_, err = database.DB.Exec(
		"INSERT INTO submissions(slug, code_hash, language, created_at) VALUES (?, ?, ?, ?)",
		submission.Slug,
		codeHash,
		submission.Language,
		submission.Timestamp,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert submission"})
		return
	}

	// Check if problem exists
	var timesSolved int
	err = database.DB.QueryRow(
		"SELECT times_solved FROM problems WHERE slug = ?",
		submission.Slug,
	).Scan(&timesSolved)

	if err == sql.ErrNoRows {
		// First time solving
		_, err = database.DB.Exec(
			"INSERT INTO problems(slug, title, difficulty, times_solved, first_solved, last_solved) VALUES (?, ?, ?, ?, ?, ?)",
			submission.Slug,
			submission.Title,
			submission.Difficulty,
			1,
			submission.Timestamp,
			submission.Timestamp,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert problem"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":      "new_problem",
			"solve_count": 1,
		})
		return
	}

	// Revisit
	newCount := timesSolved + 1

	_, err = database.DB.Exec(
		"UPDATE problems SET times_solved = ?, last_solved = ? WHERE slug = ?",
		newCount,
		submission.Timestamp,
		submission.Slug,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update problem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "revisit",
		"solve_count": newCount,
	})
}
