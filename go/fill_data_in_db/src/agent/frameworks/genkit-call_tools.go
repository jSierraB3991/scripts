package frameworks

import (
	"fmt"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/jsierra/fill_in_data/tools"
)

func resultTool(resultToolResponse string) (string, error) {
	if strings.HasPrefix(strings.ToLower(resultToolResponse), "error") {
		fmt.Println(resultToolResponse)
		return "", fmt.Errorf("%s", resultToolResponse)
	}
	return resultToolResponse, nil
}

type GetTableInput struct{}

type GetTableColumnsInput struct {
	TableName string `json:"table_name" jsonschema_description:"Nombre completo de la tabla en formato schema.tabla, por ejemplo erp.products"`
}

type ExecuteSelectInput struct {
	Query string `json:"query" jsonschema_description:"Consulta SQL SELECT que se ejecutará."`
}

type InsertDataInput struct {
	Insert string `json:"insert" jsonschema_description:"Sentencia INSERT parcial. Ejemplo: INSERT INTO erp.clientes (nombre, email)"`
	Values string `json:"values" jsonschema_description:"Valores a insertar. Ejemplo: ('Juan Perez', 'juan@mail.com')"`
}

func getToolsGenkit(g *genkit.Genkit) []ai.ToolRef {
	return []ai.ToolRef{
		genkit.DefineTool(
			g,
			"get_tables",
			"Obtiene la lista de tablas existentes en el esquema erp de PostgreSQL. Úsala para descubrir las tablas disponibles antes de consultar información de la base de datos.",
			func(ctx *ai.ToolContext, input GetTableInput) (string, error) {
				return resultTool(tools.GetTablesTool())
			},
		),

		genkit.DefineTool(
			g,
			"get_table_columns",
			"Obtiene los nombres y tipos de todas las columnas de una tabla PostgreSQL. Recibe el nombre completo en formato schema.tabla, por ejemplo erp.products. Usa esta herramienta antes de consultar datos de una tabla desconocida.",
			func(ctx *ai.ToolContext, input GetTableColumnsInput) (string, error) {
				return resultTool(tools.GetTableColumns(input.TableName))
			},
		),

		genkit.DefineTool(
			g,
			"execute_select",
			"Ejecuta una consulta SQL SELECT en PostgreSQL y devuelve los resultados en JSON. Utiliza únicamente consultas SELECT.",
			func(ctx *ai.ToolContext, input ExecuteSelectInput) (string, error) {
				return resultTool(tools.ExecuteSelect(input.Query))
			},
		),

		genkit.DefineTool(
			g,
			"insert_data",
			"Inserta datos en una tabla del esquema erp de PostgreSQL. Recibe la sentencia INSERT (sin la parte de VALUES) y los valores por separado.",
			func(ctx *ai.ToolContext, input InsertDataInput) (string, error) {
				return resultTool(tools.InsertDataTool(input.Insert, input.Values))
			},
		),
	}
}
