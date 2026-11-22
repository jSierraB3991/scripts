package ollamaagent

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ollama/ollama/api"
)

func (agent *OllamaAgent) RunChat(input string) {
	if strings.ToLower(input) == "clear memory" || strings.ToLower(input) == "reiniciar memoria" {
		agent.resetMemory()
		return
	}
	agent.messages = append(agent.messages, api.Message{
		Role:    "user",
		Content: input,
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &api.ChatRequest{
			Model:    agent.model,
			Messages: agent.messages,
			Tools:    GetTools(),
			Stream:   new(false),
		}

		var response api.ChatResponse

		err = client.Chat(ctx, req, func(resp api.ChatResponse) error {
			response = resp
			return nil
		})
		if err != nil {
			if err.Error() == "Post \"http://127.0.0.1:11434/api/chat\": context canceled" {
				return
			}
			fmt.Printf("Error al comunicar con Ollama (%s): %v\n", agent.model, err)
			return
		}

		if len(response.Message.ToolCalls) == 0 {
			if !response.Done || (response.Message.Thinking != "" && response.Message.Content == "") {
				agent.messages = append(agent.messages, response.Message)
				fmt.Printf("\n\033[1;32mThinking %s\033[0m: ", agent.model)
				fmt.Printf("%s\n", response.Message.Thinking)
				continue
			} else {
				agent.messages = append(agent.messages, response.Message)
				fmt.Printf("\033[36m%s - %s\033[0m: ", agent.GetName(), agent.model)
				fmt.Printf("%s\n", response.Message.Content)
				break
			}
		} else {
			if response.Message.Thinking != "" {
				fmt.Printf("\n\033[1;32mThinking %s\033[0m: %s\n", agent.model, response.Message.Thinking)
			}
			agent.messages = append(agent.messages, response.Message)
			agent.callTools(response)
		}
	}
}
