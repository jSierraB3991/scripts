package tools

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetTablesTool() string {
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

	rows, err := conn.Query(ctx, `SELECT table_name FROM information_schema.tables
		WHERE table_schema = $1 AND table_type = 'BASE TABLE' ORDER BY table_name`, "erp")
	if err != nil {
		return err.Error()
	}
	defer rows.Close()

	var result string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return err.Error()
		}

		result += fmt.Sprintf("erp.%s\n", table)
	}

	if err := rows.Err(); err != nil {
		return err.Error()
	}
	return result
}
