package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main5() {
	llm, err := ollama.New(ollama.WithModel("llama3.1:8b"))
	if err != nil {
		log.Fatal(err)
	}

	//query = "What is the L2 Lagrange point?. Anser me in spanish"
	query := "very briefly, tell me the difference between a comet and meteor. Anser me in spanish"
	ctx := context.Background()
	completation, err := llms.GenerateFromSinglePrompt(ctx, llm, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %s\n", completation)

	texts := []string{"meteor", "comet", "puppy"}
	embs, err := llm.CreateEmbedding(ctx, texts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got %d embedings: \n", len(embs))
	for i, v := range embs {
		fmt.Printf("%d: len=%d first-few=%v \n", i, len(v), v[:4])
	}
}
