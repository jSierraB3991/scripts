package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

// ExecuteAnonymizedSelect ejecuta una consulta SELECT permitida y anonimiza los resultados.
func ExecuteAnonymizedSelect(query string) string {
	fmt.Printf("[ExecuteAnonymizedSelect] query: %s\n", query)

	query = strings.TrimSpace(query)
	query = strings.TrimSuffix(query, ";")
	query = strings.TrimSpace(query)

	upper := strings.ToUpper(query)
	if !strings.HasPrefix(upper, "SELECT ") &&
		!strings.HasPrefix(upper, "SELECT\n") &&
		!strings.HasPrefix(upper, "SELECT\t") {
		return "Error: solo se permiten consultas SELECT"
	}

	if strings.Contains(query, ";") {
		return "Error: no se permiten múltiples sentencias SQL"
	}

	pgUrl := os.Getenv("PG_URL")
	if pgUrl == "" {
		return "Error: PG_URL no está configurada en las variables de entorno"
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, pgUrl)
	if err != nil {
		return fmt.Sprintf("Error de conexión: %v", err)
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return fmt.Sprintf("Error en la consulta: %v", err)
	}
	defer rows.Close()

	fields := rows.FieldDescriptions()
	result := make([]map[string]any, 0)

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return fmt.Sprintf("Error leyendo fila: %v", err)
		}

		row := make(map[string]any, len(fields))
		for i, field := range fields {
			row[string(field.Name)] = values[i]
		}

		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return fmt.Sprintf("Error en iteración: %v", err)
	}

	// Aplicar anonimización a todos los registros obtenidos
	anonymizedResult := AnonymizeRows(result)

	data, err := json.Marshal(anonymizedResult)
	if err != nil {
		return fmt.Sprintf("Error serializando JSON: %v", err)
	}

	return string(data)
}
