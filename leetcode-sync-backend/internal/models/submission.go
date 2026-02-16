package models

type Submission struct {
	Slug            string `json:"slug" binding:"required"`
	Title           string `json:"title" binding:"required"`
	Difficulty      string `json:"difficulty" binding:"required"`
	DescriptionHTML string `json:"description_html" binding:"required"`
	Language        string `json:"language" binding:"required"`
	Timestamp       string `json:"timestamp" binding:"required"`
	Code            string `json:"code" binding:"required"`
}
