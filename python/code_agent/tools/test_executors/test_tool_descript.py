from tools.test_executors import test_go_tool, test_python_tool
test_tools = [
    {
        "type": "function",
        "function":{
            "name": "run_go_tests",
            "description": "Corre los test del lenguage go.",
            "parameters":{
                "type": "object",
                "properties": {
                    "path": {
                        "type": "string",
                        "description": "Ruta del directorio donde están los test de golang"
                    }
                },
                "required": ["path"]
            }
        }
    },
    {
        "type": "function",
        "function": {
            "name": "run_python_tests",
            "description": "Corre los test del lenguage python.",
            "parameters":{
                "type": "object",
                "properties": {
                    "path": {
                        "type": "string",
                        "description": "Ruta del directorio donde están los test de python"
                    },
                },
                "required": ["path"]
            }
        }
    }
]

test_available_tools = {
    "run_go_tests": test_go_tool.run_go_tests,
    "run_python_tests": test_python_tool.run_python_tests
}

