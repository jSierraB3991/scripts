from tools.database import save_memory_tool, update_memory_tool, get_memory_by_key_tool
database_tools = [
    {
        "type": "function",
        "function":{
            "name": "save_memory",
            "description": "Guarda en memoria los datos que debe recordar el agente, trata de guardar cuando necesites usar una tool, para no perder el flujo",
            "parameters":{
                "type": "object",
                "properties": {
                    "key": {
                        "type": "string",
                        "description": "llave generica para poder identificar el dato guardado, si das uno repetido dará error, siempre debe iniciar con 'key_'"
                    },
                    "role_agent": {
                        "type": "string",
                        "description": "tú siempre deberias llamar a este parametro con el valor 'agent'",
                    },
                    "content": {
                        "type": "string",
                        "description": "dato que el agente considere importante para guardar, tratar de que sean bastante fáciles de leer, como un string, o un json"
                    },
                    "description": {
                        "type": "string",
                        "description": "valor que hace que el agente pueda tener un poco de contexto del valor guardado, por lo tanto es importante, aunque no requerido"
                    }
                },
                "required": ["key", "content", "role_agent"]
            }
        }
    },
    {
        "type": "function",
        "function":{
            "name": "update_memory",
            "description": "Actualiza el contenido de una memoria en la base de datos por la key",
            "parameters":{
                "type": "object",
                "properties": {
                    "key": {
                        "type": "string",
                        "description": "Llave necesaria para poder actualizar la memoria"
                    },
                    "content": {
                        "type": "string",
                        "description": "contenido que va a sobreescribir el anterior, tratar de que sean bastante fáciles de leer, como un string, o un json"
                    }
                },
                "required": ["key", "content"]
            }
        }
    },
    {
        "type": "function",
        "function": {
            "name": "get_memory_by_key_tool",
            "description": "Busca en la base de datos, por la key, si esa memoria ya está guardada",
            "parameters":{
                "type": "object",
                "properties": {
                    "key": {
                        "type": "string",
                        "description": "key de la base de datos, para acceder rapidamente a un dato de la memoria"
                    },
                },
                "required": ["key"]
            }
        }
    }
]

database_available_tools = {
    "save_memory": save_memory_tool.save_memory,
    "update_memory": update_memory_tool.update_memory,
    "get_memory_by_key_tool": get_memory_by_key_tool.get_memory_by_key_tool
}

