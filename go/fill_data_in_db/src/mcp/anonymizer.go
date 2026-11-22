package mcp

import (
	"fmt"
	"strings"
)

// AnonymizeValue aplica reglas de anonimización según el nombre de la columna y el valor.
func AnonymizeValue(columnName string, val any) any {
	if val == nil {
		return nil
	}

	colLower := strings.ToLower(strings.TrimSpace(columnName))
	strVal := fmt.Sprintf("%v", val)

	switch {
	case strings.Contains(colLower, "password") || strings.Contains(colLower, "hash"):
		return "********"

	case strings.Contains(colLower, "email"):
		parts := strings.Split(strVal, "@")
		if len(parts) == 2 {
			name := parts[0]
			if len(name) <= 2 {
				return "***@" + parts[1]
			}
			return name[:2] + "***@" + parts[1]
		}
		return "user@example.com"

	case strings.Contains(colLower, "phone") || strings.Contains(colLower, "document_number"):
		if len(strVal) <= 4 {
		return "****"
		}
		return strings.Repeat("*", len(strVal)-4) + strVal[len(strVal)-4:]

	case strings.Contains(colLower, "address"):
		return "Dirección Anonimizada"

	case strings.Contains(colLower, "name") || strings.Contains(colLower, "description") || strings.Contains(colLower, "notes") || strings.Contains(colLower, "reference"):
		// Anonimizar nombres o textos largos manteniendo una estructura genérica
		if len(strVal) > 0 {
			return fmt.Sprintf("Anonymized_%s", colLower)
		}
		return ""

	default:
		return val
	}
}

// AnonymizeRow toma un mapa de una fila (columna -> valor) y anonimiza los campos sensibles.
func AnonymizeRow(row map[string]any) map[string]any {
	anonymized := make(map[string]any, len(row))
	for k, v := range row {
		anonymized[k] = AnonymizeValue(k, v)
	}
	return anonymized
}

// AnonymizeRows anonimiza un slice de filas.
func AnonymizeRows(rows []map[string]any) []map[string]any {
	result := make([]map[string]any, len(rows))
	for i, row := range rows {
		result[i] = AnonymizeRow(row)
	}
	return result
}
