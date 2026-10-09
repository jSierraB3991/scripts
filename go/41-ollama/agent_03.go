package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
)

type ReadFile struct {
	PathRouteFile string `json:"path_route_file" jsonschema:"description=Ruta del archivo a leer"`
}

func main3() {
	ctx := context.Background()

	o := &ollama.Ollama{
		ServerAddress: "http://localhost:11434", // Default Ollama server
		Timeout:       600,                      // Response timeout in seconds
	}

	g := genkit.Init(ctx, genkit.WithPlugins(o))
	readFileTool := genkit.DefineTool(g, "readFileTool",
		"Return the content by define route in routeFile file.",
		func(toolCtx *ai.ToolContext, routeFile ReadFile) (string, error) {
			log.Printf("call tool %s with params %v\n", "readFileTool", routeFile)
			baseDir := "."
			// 2. Construir la ruta final del archivo
			pathRouteFile := routeFile.PathRouteFile
			if pathRouteFile == "" {
				return "", fmt.Errorf("Error el nombre del archivo es obligatorio")
			}
			filePatAbs, err := filepath.Abs(filepath.Join(baseDir, pathRouteFile))
			if pathRouteFile == "" {
				return "", fmt.Errorf("Error obteniendo la ruta abosluta filepath.Abs(%s + '/' + %s)", baseDir, pathRouteFile)
			}

			baseDirAbs, err := filepath.Abs(baseDir)
			if pathRouteFile == "" {
				return "", fmt.Errorf("Error obteniendo la ruta abosluta filepath.Abs(baseDir)")
			}

			// 3. Sanitización (CWE-22): Verificar que filePatAbs esté dentro de baseAbs
			rel, err := filepath.Rel(baseDirAbs, filePatAbs)
			if err != nil || strings.HasPrefix(rel, "..") || rel == ".." {
				return "Error: Intento de acceso denegado fuera del directorio permitido", err
			}

			// 4. Verificar si la ruta es una carpeta o un archivo
			info, err := os.Stat(filePatAbs)
			if err != nil {
				return fmt.Sprintf("Error: No se encontró el archivo en %s", baseDir), err
			}

			if info.IsDir() {
				return "", fmt.Errorf("Error: En la carpeta %s solo encuentro un %s y es una carpeta", baseDir, pathRouteFile)
			}

			// 5. Leer el contenido del archivo
			// #nosec G304 -- Sanitización validada mediante filepath.Rel arriba
			contenido, err := os.ReadFile(filePatAbs)
			if err != nil {
				return fmt.Sprintf("Error al leer el archivo %s", pathRouteFile), err
			}

			return string(contenido), nil
		},
	)

	// Any model installed on the server resolves by name.
	resp, err := genkit.Generate(ctx, g,
		ai.WithPrompt("Lee el archivo agent_03.go y describelo para mi"),
		ai.WithModelName("ollama/llama3.1:8b"),
		ai.WithTools(readFileTool),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.Text())
	fmt.Println(resp)

}
