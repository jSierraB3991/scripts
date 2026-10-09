package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/memory"
)

func main6() {
	openModel := "llama3.1:8b"
	ctx := context.Background()
	llm, err := ollama.New(ollama.WithModel(openModel))
	if err != nil {
		log.Fatal(err)
	}

	bufferMemory := memory.NewConversationBuffer()
	conversationChain := chains.NewConversation(llm, bufferMemory)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("💬 You: ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading user input %s\n", err)
			continue
		}

		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			fmt.Println("⚠️ Please enter a message")
			continue
		}

		if strings.ToLower(userInput) == "salir" || strings.ToLower(userInput) == "exit" {
			fmt.Println("Adios")
			break
		}

		fmt.Printf("🤖 %s ", openModel)
		fmt.Print("Thinking")

		response, err := chains.Run(ctx, conversationChain, userInput)
		fmt.Printf("\n🤖 %s: ", openModel)
		if err != nil {
			fmt.Printf("❌ Error %s \n", err)
			continue
		}

		fmt.Println(response)
	}
}
