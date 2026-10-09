package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

func main1() {
	ctx := context.Background()
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	req := &api.GenerateRequest{
		Model:  "llama3.1:8b",
		Prompt: "Write a short poem about Go programming.",
		Stream: nil, // default streaming
	}

	err = client.Generate(ctx, req, func(resp api.GenerateResponse) error {
		fmt.Print(resp.Response)
		if resp.Done {
			fmt.Println("\n[Generation complete]")
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
