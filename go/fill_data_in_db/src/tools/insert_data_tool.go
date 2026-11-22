package tools

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
)

func InsertDataTool(insert string, values string) string {
	pgUrl := os.Getenv("PG_URL")
	if pgUrl == "" {
		return "Error: PG_URL no está configurada en las variables de entorno"
	}

	fmt.Printf("[InsertDataTool] params [insert]: %s  [values] %s\n", insert, values)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, pgUrl)
	if err != nil {
		return err.Error()
	}
	defer conn.Close(ctx)

	upperInsert := strings.ToUpper(strings.TrimSpace(insert))

	// Solo permitir INSERT
	if !strings.HasPrefix(upperInsert, "INSERT INTO") {
		return "Error: solo se permiten sentencias INSERT INTO"
	}

	// Validar que no haya punto y coma (evitar múltiples statements)
	if strings.Contains(insert, ";") || strings.Contains(values, ";") {
		return "Error: no se permiten múltiples sentencias"
	}

	var query string
	if strings.Contains(upperInsert, "VALUES") {
		query = fmt.Sprintf("%s %s", insert, values)
	} else {
		query = fmt.Sprintf("%s VALUES %s", insert, values)
	}

	tag, err := conn.Exec(ctx, query)
	if err != nil {
		return fmt.Sprintf("Error al insertar: %v", err)
	}

	return fmt.Sprintf("Insert exitoso. Filas afectadas: %d", tag.RowsAffected())
}
