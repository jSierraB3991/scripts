package tools

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func GetTableColumns(tableName string) string {
	fmt.Printf("[GetTableColumns] params [tableName]: %s\n", tableName)
	cleanName := strings.TrimSpace(tableName)
	parts := strings.Split(cleanName, ".")

	var schema, table string
	if len(parts) == 1 && cleanName != "" {
		schema, table = "erp", cleanName
	} else if len(parts) == 2 {
		schema, table = parts[0], parts[1]
	} else {
		return "Error: formato de tabla inválido. Usa 'tabla' o 'schema.tabla'"
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

	rows, err := conn.Query(ctx, `
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_schema = $1
		  AND table_name = $2
		ORDER BY ordinal_position
	`, schema, table)
	if err != nil {
		return err.Error()
	}
	defer rows.Close()

	var result strings.Builder

	for rows.Next() {
		var column, dataType string

		if err := rows.Scan(&column, &dataType); err != nil {
			return err.Error()
		}

		fmt.Fprintf(&result, "%s: %s\n", column, dataType)
	}

	if err := rows.Err(); err != nil {
		return err.Error()
	}

	if result.Len() == 0 {
		return fmt.Sprintf("Error: No existe la tabla %s.%s o no tiene columnas visibles", schema, table)
	}

	return result.String()
}
