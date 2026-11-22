package ollamaagent

import (
	"fmt"
	"os"

	"github.com/ollama/ollama/api"
)

type OllamaAgent struct {
	model        string
	systemPrompt string
	messages     []api.Message
}

func NewOllamaAgent(prompt string) *OllamaAgent {
	model := os.Getenv("OLLAMA_MODEL")
	if model == "" {
		//model = "nemotron-3.5-bolt:e4b"
		model = "qwen3.5:4b"
	}
	return &OllamaAgent{
		model:        model,
		systemPrompt: prompt,
		messages: []api.Message{
			{
				Role:    "system",
				Content: prompt,
			},
		},
	}
}

func (agent *OllamaAgent) GetName() string {
	return "Ollama"
}

func (agent *OllamaAgent) resetMemory() {
	agent.messages = []api.Message{
		{
			Role:    "system",
			Content: agent.systemPrompt,
		},
	}
	fmt.Printf("🧹 Memoria de Ollama vaciada (historial reiniciado y modelo '%s' liberado de VRAM/RAM).\n", agent.model)
}
