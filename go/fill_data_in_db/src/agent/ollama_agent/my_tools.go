package ollamaagent

import (
	"github.com/ollama/ollama/api"
)

func GetTools() api.Tools {
	return api.Tools{
		{
			Type: "function",
			Function: api.ToolFunction{
				Name:        "get_tables",
				Description: "Obtiene la lista de tablas existentes en el esquema erp de PostgreSQL. Úsala para descubrir las tablas disponibles antes de consultar información de la base de datos.",
				Parameters: api.ToolFunctionParameters{
					Type:       "object",
					Properties: api.NewToolPropertiesMap(),
				},
			},
		},
		{
			Type: "function",
			Function: api.ToolFunction{
				Name:        "get_table_columns",
				Description: "Obtiene los nombres y tipos de todas las columnas de una tabla PostgreSQL. Recibe el nombre completo en formato schema.tabla, por ejemplo erp.products. Usa esta herramienta antes de consultar datos de una tabla desconocida.",
				Parameters: api.ToolFunctionParameters{
					Type: "object",
					Properties: func() *api.ToolPropertiesMap {
						props := api.NewToolPropertiesMap()
						props.Set("table_name", api.ToolProperty{
							Type:        api.PropertyType{"string"},
							Description: "Nombre completo de la tabla en formato schema.tabla, por ejemplo erp.products",
						})
						return props
					}(),
					Required: []string{"table_name"},
				},
			},
		}, {
			Type: "function",
			Function: api.ToolFunction{
				Name:        "execute_select",
				Description: "Ejecuta una consulta SQL SELECT en PostgreSQL y devuelve los resultados en JSON. Utiliza únicamente consultas SELECT.",
				Parameters: api.ToolFunctionParameters{
					Type: "object",
					Properties: func() *api.ToolPropertiesMap {
						props := api.NewToolPropertiesMap()
						props.Set("query", api.ToolProperty{
							Type:        api.PropertyType{"string"},
							Description: "Consulta SQL SELECT que se ejecutará.",
						})
						return props
					}(),
					Required: []string{"query"},
				},
			},
		},
		{
			Type: "function",
			Function: api.ToolFunction{
				Name:        "insert_data",
				Description: "Inserta datos en una tabla del esquema erp de PostgreSQL. Recibe la sentencia INSERT (sin la parte de VALUES) y los valores por separado.",
				Parameters: api.ToolFunctionParameters{
					Type:       "object",
					Properties: GetInsertDataProps(),
					Required:   []string{"insert", "values"},
				},
			},
		},
	}
}

func GetInsertDataProps() *api.ToolPropertiesMap {
	props := api.NewToolPropertiesMap()
	props.Set("insert", api.ToolProperty{
		Type:        api.PropertyType{"string"},
		Description: "Sentencia INSERT parcial. Ejemplo: INSERT INTO erp.clientes (nombre, email)",
	})
	props.Set("values", api.ToolProperty{
		Type:        api.PropertyType{"string"},
		Description: "Valores a insertar. Ejemplo: ('Juan Perez', 'juan@mail.com')",
	})
	return props
}
