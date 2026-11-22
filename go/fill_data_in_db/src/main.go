package main

import (
	"flag"

	"github.com/joho/godotenv"
	"github.com/jsierra/fill_in_data/agent"
)

func main() {
	_ = godotenv.Load()
	modelParam := flag.String("llm", "ollama", "Llm a usar: ollama, gemini (genai), open-router (openrouter), genkit (genkit ollama), langchain (langchaing ollama)")
	flag.Parse()
	agent.Start(*modelParam)
}
