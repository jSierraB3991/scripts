package ollamaagent

import (
	"fmt"
	"strings"

	"github.com/jsierra/fill_in_data/tools"
	"github.com/ollama/ollama/api"
)

func (agent *OllamaAgent) insertData(call api.ToolCall) (string, string) {
	insertValue, ok := call.Function.Arguments.Get("insert")
	if !ok {
		return "", "ERROR: falta el argumento insert"
	}

	insert, ok := insertValue.(string)
	if !ok {
		return "", "ERROR: insert debe ser un string"
	}

	valuesValue, ok := call.Function.Arguments.Get("values")
	if !ok {
		return "", "ERROR: falta el argumento values"
	}

	values, ok := valuesValue.(string)
	if !ok {
		return "", "ERROR: values debe ser un string"
	}
	var query string
	if strings.Contains(insert, "VALUES") {
		query = fmt.Sprintf("%s %s", insert, values)
	} else {
		query = fmt.Sprintf("%s VALUES %s", insert, values)
	}

	return query, tools.InsertDataTool(insert, values)
}

func (agent *OllamaAgent) executeSelect(call api.ToolCall) (string, string) {
	value, ok := call.Function.Arguments.Get("query")
	if !ok {
		return "", "ERROR: falta el argumento query"
	}

	query, ok := value.(string)
	if !ok {
		return "", "ERROR: query debe ser un string"
	}

	return query, tools.ExecuteSelect(query)
}

func (agent *OllamaAgent) getTablesColumns(call api.ToolCall) (string, string) {
	value, ok := call.Function.Arguments.Get("table_name")
	if !ok {
		return "", "ERROR: falta el argumento table_name"
	}

	tableName, ok := value.(string)
	if !ok {
		return "", "ERROR: table_name debe ser un string"
	}

	return tableName, tools.GetTableColumns(tableName)
}

func (agent *OllamaAgent) callToolGetTables() string {
	return tools.GetTablesTool()
}

func (agent *OllamaAgent) callTools(response api.ChatResponse) {
	for _, call := range response.Message.ToolCalls {
		var result string
		switch call.Function.Name {
		case "get_tables":
			fmt.Printf("🤖 Ollama requested function call: %s\n", call.Function.Name)
			result = agent.callToolGetTables()
		case "get_table_columns":
			fmt.Printf("🤖 Ollama requested function call: %s\n", call.Function.Name)
			_, result = agent.getTablesColumns(call)
		case "execute_select":
			fmt.Printf("🤖 Ollama requested function call: %s\n", call.Function.Name)
			_, result = agent.executeSelect(call)
		case "insert_data":
			fmt.Printf("🤖 Ollama requested function call: %s\n", call.Function.Name)
			_, result = agent.insertData(call)
		default:
			result = fmt.Sprintf("Error: tool desconocida: %s\n", call.Function.Name)
			fmt.Print(result)
		}

		agent.messages = append(agent.messages, api.Message{
			Role:       "tool",
			ToolName:   call.Function.Name,
			ToolCallID: call.ID,
			Content:    result,
		})
	}
}
