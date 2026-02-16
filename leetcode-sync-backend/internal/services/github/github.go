package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type GitHubService struct {
	Token  string
	Owner  string
	Repo   string
	Branch string
}

func NewGitHubService() *GitHubService {
	fmt.Println("DEBUG: GITHUB_TOKEN=", os.Getenv("GITHUB_TOKEN"))
	fmt.Println("DEBUG: GITHUB_OWNER=", os.Getenv("GITHUB_OWNER"))
	fmt.Println("DEBUG: GITHUB_REPO=", os.Getenv("GITHUB_REPO"))
	fmt.Println("DEBUG: GITHUB_BRANCH=", os.Getenv("GITHUB_BRANCH"))

	return &GitHubService{
		Token:  os.Getenv("GITHUB_TOKEN"),
		Owner:  os.Getenv("GITHUB_OWNER"),
		Repo:   os.Getenv("GITHUB_REPO"),
		Branch: os.Getenv("GITHUB_BRANCH"),
	}
}

func (g *GitHubService) CreateOrUpdateFile(path string, content string, message string) error {
	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/contents/%s",
		g.Owner,
		g.Repo,
		path,
	)

	encoded := base64.StdEncoding.EncodeToString([]byte(content))

	body := map[string]interface{}{
		"message": message,
		"content": encoded,
		"branch":  g.Branch,
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", "Bearer "+g.Token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("GitHub API error: %d", resp.StatusCode)
	}

	return nil
}
