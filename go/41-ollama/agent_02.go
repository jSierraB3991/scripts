package main

import (
	"context"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
)

func main2() {
	ctx := context.Background()

	o := &ollama.Ollama{
		ServerAddress: "http://localhost:11434", // Default Ollama server
		Timeout:       600,                      // Response timeout in seconds
	}

	g := genkit.Init(ctx, genkit.WithPlugins(o))

	// Any model installed on the server resolves by name.
	resp, err := genkit.Generate(ctx, g,
		ai.WithPrompt("Why sky is blue?, answerme in spanish"),
		ai.WithModelName("ollama/llama3.1:8b"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Acceder a las métricas de uso de la respuesta
	if resp.Usage != nil {
		fmt.Printf("Tokens de Entrada: %d\n", resp.Usage.InputTokens)
		fmt.Printf("Tokens de Salida: %d\n", resp.Usage.OutputTokens)
		fmt.Printf("Caracteres de Entrada: %d\n", resp.Usage.InputCharacters)
		fmt.Printf("Caracteres de Salida: %d\n", resp.Usage.OutputCharacters)
		fmt.Printf("Tokens Totales: %d\n", resp.Usage.TotalTokens)

		// Aquí puedes guardar este consumo en tu base de datos (por ejemplo, por ID de usuario)
		// y bloquear temporalmente sus solicitudes si excede una cuota establecida por ti.
	}

	log.Println(resp.Text())
}
