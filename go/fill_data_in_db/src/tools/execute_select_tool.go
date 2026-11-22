package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func ExecuteSelect(query string) string {
	fmt.Printf("[ExecuteSelect] params [query]: %s\n", query)

	query = strings.TrimSpace(query)
	query = strings.TrimSuffix(query, ";")
	query = strings.TrimSpace(query)

	upper := strings.ToUpper(query)
	if !strings.HasPrefix(upper, "SELECT ") &&
		!strings.HasPrefix(upper, "SELECT\n") &&
		!strings.HasPrefix(upper, "SELECT\t") {
		return "Error: solo se permiten consultas SELECT"
	}

	// Rechazar sentencias múltiples.
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
		return err.Error()
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return err.Error()
	}
	defer rows.Close()

	fields := rows.FieldDescriptions()
	result := make([]map[string]any, 0)

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return err.Error()
		}

		row := make(map[string]any, len(fields))
		for i, field := range fields {
			row[string(field.Name)] = values[i]
		}

		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return err.Error()
	}

	data, err := json.Marshal(result)
	if err != nil {
		return err.Error()
	}

	return string(data)
}
