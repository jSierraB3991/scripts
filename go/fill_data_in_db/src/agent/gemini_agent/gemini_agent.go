package geminiagent

import (
	"os"

	"google.golang.org/genai"
)

type GeminiAgent struct {
	model        string
	systemPrompt string
	messages     []*genai.Content
}

func NewGeminiAgent(prompt string) *GeminiAgent {
	model := os.Getenv("GEMINI_MODEL")
	if model == "" {
		model = "gemini-3.5-flash-lite"
	}
	return &GeminiAgent{
		model:        model,
		systemPrompt: prompt,
		messages:     []*genai.Content{},
	}
}

func (agent *GeminiAgent) GetName() string {
	return "Gemini"
}
