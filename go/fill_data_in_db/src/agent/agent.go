package agent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	geminiagent "github.com/jsierra/fill_in_data/agent/gemini_agent"
	ollamaagent "github.com/jsierra/fill_in_data/agent/ollama_agent"
)

type AgentInterface interface {
	GetName() string
	RunChat(input string)
}

func otherCommands(input string) bool {
	switch input {
	case "":
		return true
	case "clear", "cls":
		fmt.Print("\033[H\033[2J")
		return true
	default:
		return false
	}
}

func Start(llm string) {

	prompt := "Eres un agente especializado en consultar una base de datos PostgreSQL." +
		"El esquema de trabajo es erp." +
		"Reglas:" +
		"1. A menos que el usuario lo diga explicitamentente, no hagas inserts ni updates y mucho menos delete" +
		"2. Antes de consultar datos de una tabla desconocida, descubre las tablas disponibles." +
		"3. Usa information_schema.columns para conocer las columnas de cada tabla." +
		"4. Consulta únicamente tablas del esquema erp." +
		"5. Utiliza exclusivamente consultas de lectura." +
		"6. No inventes nombres de tablas, columnas ni resultados." +
		"7. Si una consulta falla, analiza el error y corrige la consulta." +
		"8. Responde basándote únicamente en los resultados obtenidos." +
		"9. Si necesitas varias consultas, ejecuta las herramientas necesarias antes de responder.." +
		"10. los campos uuid, debes buscarlos así SELECT id::text"

	selectedLLM := strings.ToLower(strings.TrimSpace(llm))
	switch selectedLLM {
	case "genai":
		selectedLLM = "gemini"
	case "openrouter":
		selectedLLM = "open-router"
	}

	llms := map[string]AgentInterface{
		"ollama": ollamaagent.NewOllamaAgent(prompt),
		"gemini": geminiagent.NewGeminiAgent(prompt),
		//"open-router": openrouteragent.NewOpenRouterAgent(prompt),
	}
	selectedAgent := llms[selectedLLM]
	if selectedAgent == nil {
		log.Fatalf("el llm '%s' no está configurado (opciones válidas: ollama, gemini, open-router)", llm)
	}

	fmt.Println("Escribe salir/exit para terminar.")
	fmt.Printf("Chat con %s\n", selectedAgent.GetName())

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n\033[32mTú\033[0m: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		trimmedInput := strings.TrimSpace(input)
		lower := strings.ToLower(trimmedInput)
		if lower == "exit" || lower == "salir" {
			break
		}

		if otherCommands(lower) {
			continue
		}

		selectedAgent.RunChat(trimmedInput)
	}
}
