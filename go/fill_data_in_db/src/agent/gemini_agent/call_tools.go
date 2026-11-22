package geminiagent

import (
	"fmt"
	"strings"

	"github.com/jsierra/fill_in_data/tools"
	"google.golang.org/genai"
)

func resultTool(resultToolResponse string) (string, error) {
	if strings.HasPrefix(strings.ToLower(resultToolResponse), "error") {
		fmt.Println(resultToolResponse)
		return "", fmt.Errorf("%s", resultToolResponse)
	}
	return resultToolResponse, nil
}

func (agent *GeminiAgent) insertData(call *genai.FunctionCall) (string, error) {
	insert, ok := call.Args["insert"].(string)
	if !ok {
		return "", fmt.Errorf("ERROR: falta el argumento insert")
	}

	values, ok := call.Args["values"].(string)
	if !ok {
		return "", fmt.Errorf("ERROR: falta el argumento values")
	}
	return resultTool(tools.InsertDataTool(insert, values))
}

func (agent *GeminiAgent) executeSelect(call *genai.FunctionCall) (string, error) {
	query, ok := call.Args["query"].(string)
	if !ok {
		return "", fmt.Errorf("ERROR: falta el argumento query")
	}

	return resultTool(tools.ExecuteSelect(query))
}

func (agent *GeminiAgent) getTablesColumns(call *genai.FunctionCall) (string, error) {
	tableName, ok := call.Args["table_name"].(string)
	if !ok {
		return "", fmt.Errorf("ERROR: falta el argumento table_name")
	}
	return resultTool(tools.GetTableColumns(tableName))
}

func (agent *GeminiAgent) callToolGetTables() (string, error) {
	return resultTool(tools.GetTablesTool())
}

func (agent *GeminiAgent) callTools(functionCalls []*genai.FunctionCall) {
	var parts []*genai.Part
	for _, call := range functionCalls {
		fmt.Printf("🤖 Gemini requested function call: %s with args: %v\n", call.Name, call.Args)
		var result string
		var err error
		switch call.Name {
		case "get_tables":
			result, err = agent.callToolGetTables()
		case "get_table_columns":
			result, err = agent.getTablesColumns(call)
		case "execute_select":
			result, err = agent.executeSelect(call)

		case "insert_data":
			result, err = agent.insertData(call)
		default:
			result, err = "", fmt.Errorf("ERROR, invalid name tools %s", call.Name)
		}

		var mapResult map[string]any
		if err != nil {
			mapResult = map[string]any{"error": err.Error(), "tool": call.Name, "args": call.Args}
		} else {
			mapResult = map[string]any{"result": result, "tool": call.Name, "args": call.Args}
		}

		parts = append(parts, genai.NewPartFromFunctionResponse(call.Name, mapResult))
	}

	agent.messages = append(agent.messages, &genai.Content{
		Role:  "user",
		Parts: parts,
	})
}
