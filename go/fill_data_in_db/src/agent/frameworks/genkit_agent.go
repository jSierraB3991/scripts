package frameworks

import (
	"context"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core/api"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
)

type GenkitAgent struct {
	plugin       api.Plugin
	model        string
	systemPrompt string
	messages     map[string]any
}

func NewGenkitAgent(systemPrompt string) *GenkitAgent {
	return &GenkitAgent{
		plugin: &ollama.Ollama{
			ServerAddress: "http://localhost:11434", // Default Ollama server
			Timeout:       600,                      // Response timeout in seconds
		},
		systemPrompt: systemPrompt,
		model:        "qwen3.5:9b",
	}
}

func (agent *GenkitAgent) GetName() string {
	return "Ollama Whit Genkit"
}

func (agent *GenkitAgent) RunChat(userInput string) {
	ctx := context.Background()
	g := genkit.Init(ctx, nil)
	chatAgent, err := genkit.Generate(ctx, g,
		ai.WithSystem(agent.systemPrompt),
		ai.WithPrompt(userInput),
		ai.WithModelName(agent.model),
	)
	if err != nil {
		log.Println(err)
		return
	}
	// Acceder a las métricas de uso de la respuesta
	if chatAgent.Usage != nil {
		fmt.Printf("Tokens de Entrada: %d\n", chatAgent.Usage.InputTokens)
		fmt.Printf("Tokens de Salida: %d\n", chatAgent.Usage.OutputTokens)
		fmt.Printf("Caracteres de Entrada: %d\n", chatAgent.Usage.InputCharacters)
		fmt.Printf("Caracteres de Salida: %d\n", chatAgent.Usage.OutputCharacters)
		fmt.Printf("Tokens Totales: %d\n", chatAgent.Usage.TotalTokens)

		// Aquí puedes guardar este consumo en tu base de datos (por ejemplo, por ID de usuario)
		// y bloquear temporalmente sus solicitudes si excede una cuota establecida por ti.
	}

	log.Println(resp.Text())
}
