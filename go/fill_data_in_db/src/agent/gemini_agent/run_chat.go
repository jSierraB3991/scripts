package geminiagent

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

func GetConfigGenerator(systemPrompt string) *genai.GenerateContentConfig {
	tools := getTools()
	return &genai.GenerateContentConfig{
		Tools: []*genai.Tool{tools},
		SystemInstruction: &genai.Content{
			Parts: []*genai.Part{
				genai.NewPartFromText(systemPrompt),
			},
		},
	}
}

func (agent *GeminiAgent) RunChat(input string) {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		fmt.Printf("Error al inicializar cliente Gemini: %v\n", err)
		return
	}
	agent.messages = append(agent.messages, &genai.Content{
		Role: "user",
		Parts: []*genai.Part{
			genai.NewPartFromText(input),
		},
	})
	for {
		res, err := client.Models.GenerateContent(
			ctx,
			agent.model,
			agent.messages,
			GetConfigGenerator(agent.systemPrompt),
		)
		if err != nil {
			fmt.Printf("Error request to %s: %s\n", agent.model, err)
			break
		}
		functionCalls := res.FunctionCalls()
		if len(functionCalls) == 0 {
			if len(res.Candidates) > 0 && res.Candidates[0].Content != nil {
				agent.messages = append(agent.messages, res.Candidates[0].Content)
			}

			fmt.Printf("\033[36m%s - %s\033[0m: ", agent.GetName(), agent.model)
			fmt.Printf("%s\n", res.Text())
			break
		} else {
			if len(res.Candidates) > 0 && res.Candidates[0].Content != nil {
				agent.messages = append(agent.messages, res.Candidates[0].Content)
			}
			agent.callTools(functionCalls)
		}
	}
}
