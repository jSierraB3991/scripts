package tools

import "github.com/ollama/ollama/api"

func SumarTool(a, b float64) float64 {
	return a + b
}

func GetToolsExamples() api.Tools {
	props := api.NewToolPropertiesMap()

	props.Set("a", api.ToolProperty{
		Type:        api.PropertyType{"number"},
		Description: "Primer número",
	})

	props.Set("b", api.ToolProperty{
		Type:        api.PropertyType{"number"},
		Description: "Segundo número",
	})

	return api.Tools{
		{
			Type: "function",
			Function: api.ToolFunction{
				Name:        "SumarTool",
				Description: "Suma dos números.",
				Parameters: api.ToolFunctionParameters{
					Type:       "object",
					Properties: props,
					Required:   []string{"a", "b"},
				},
			},
		},
	}
}
