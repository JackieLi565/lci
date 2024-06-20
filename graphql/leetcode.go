package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	GraphqlEndpoint = "https://leetcode.com/graphql/"
)

const (
	Easy   Difficulty = "Easy"
	Medium Difficulty = "Medium"
	Hard   Difficulty = "Hard"
)

type Difficulty string

type QuestionResponse struct {
	Data struct {
		Question QuestionPayload `json:"question"`
	} `json:"data"`
}

type QuestionPayload struct {
	QuestionID string     `json:"questionId"`
	Title      string     `json:"title"`
	TitleSlug  string     `json:"titleSlug"`
	Difficulty Difficulty `json:"difficulty"`
}

func QuestionData(titleSlug string) (*QuestionResponse, error) {
	query := generateQuestionQuery(titleSlug)

	req, err := http.NewRequest("POST", GraphqlEndpoint, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var questionPayload *QuestionResponse
	if err := json.Unmarshal(body, &questionPayload); err != nil {
		return nil, err
	}

	return questionPayload, nil
}

func ExtractTitleSlug(url string) (string, error) {
	parts := strings.Split(url, "/")
	for i, part := range parts {
		if part == "problems" && i+1 < len(parts) {
			return parts[i+1], nil
		}
	}
	return "", fmt.Errorf("provided URL does not contain a titleSlug")
}

func generateQuestionQuery(titleSlug string) string {
	return fmt.Sprintf(`{
		"query": "query questionTitle($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    title\n    titleSlug\n    difficulty\n  }\n}\n    ",
		"variables": {
			"titleSlug": %q
		},
		"operationName": "questionTitle"
	}`, titleSlug)
}
