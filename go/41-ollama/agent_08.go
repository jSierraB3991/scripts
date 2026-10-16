package main

import (
	"context"
	"fmt"
	"log"

	agent "github.com/OpenRouterTeam/go-agent"
	"github.com/joho/godotenv"
)

type WeatherInput struct {
	Location string `json:"location" jsonschema:"required"`
}

func main() {
	fmt.Println("Star")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client := agent.NewOpenRouter(agent.OpenRouterOptions{})

	weather := agent.MustNewTool(agent.ToolConfig[WeatherInput]{
		Name:        "get_weather",
		Description: "Get the current weather for a location",
		Execute: func(ctx context.Context, in WeatherInput, tc agent.ToolExecuteContext) (any, error) {
			return map[string]any{"temperature": 72, "condition": "sunny", "location": in.Location}, nil
		},
	})

	result, err := agent.CallModel(ctx, client, agent.CallModelInput{
		Model: "nvidia/nemotron-3.5-super-vl:free",
		Input: "What is the weather in San Francisco?",
		Tools: []agent.Tool{weather},
	})
	if err != nil {
		panic(err)
	}

	text, err := result.Text(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}
