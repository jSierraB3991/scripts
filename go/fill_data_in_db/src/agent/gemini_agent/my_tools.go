package geminiagent

import "google.golang.org/genai"

func getTools() *genai.Tool {
	return &genai.Tool{
		FunctionDeclarations: []*genai.FunctionDeclaration{
			&genai.FunctionDeclaration{
				Name:        "get_tables",
				Description: "Obtiene la lista de tablas existentes en el esquema erp de PostgreSQL. Úsala para descubrir las tablas disponibles antes de consultar información de la base de datos.",
				Parameters: &genai.Schema{
					Type:       genai.TypeObject,
					Properties: nil,
				},
			},
			&genai.FunctionDeclaration{
				Name:        "get_table_columns",
				Description: "Obtiene los nombres y tipos de todas las columnas de una tabla PostgreSQL. Recibe el nombre completo en formato schema.tabla, por ejemplo erp.products. Usa esta herramienta antes de consultar datos de una tabla desconocida.",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"table_name": {
							Type:        genai.TypeString,
							Description: "Nombre completo de la tabla en formato schema.tabla, por ejemplo erp.products",
						},
					},
					Required: []string{"table_name"},
				},
			},
			&genai.FunctionDeclaration{
				Name:        "execute_select",
				Description: "Ejecuta una consulta SQL SELECT en PostgreSQL y devuelve los resultados en JSON. Utiliza únicamente consultas SELECT.",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"query": {
							Type:        genai.TypeString,
							Description: "Consulta SQL SELECT que se ejecutará.",
						},
					},
					Required: []string{"query"},
				},
			},
			&genai.FunctionDeclaration{
				Name:        "insert_data",
				Description: "Inserta datos en una tabla del esquema erp de PostgreSQL. Recibe la sentencia INSERT (sin la parte de VALUES) y los valores por separado.",
				Parameters: &genai.Schema{
					Type: genai.TypeObject,
					Properties: map[string]*genai.Schema{
						"insert": {
							Type:        genai.TypeString,
							Description: "Sentencia INSERT parcial. Ejemplo: INSERT INTO erp.clientes (nombre, email)",
						},
						"values": {
							Type:        genai.TypeString,
							Description: "Valores a insertar. Ejemplo: ('Juan Perez', 'juan@mail.com')",
						},
					},
					Required: []string{"insert", "values"},
				},
			},
		},
	}
}
