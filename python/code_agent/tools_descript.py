from tools.tool_files import *
from tools.tool_db import *
my_tools = [
    {
        "type": "function",
        "function":{
            "name": "list_directory",
            "description": "Lista los archivos y carpetas de un directorio.",
            "parameters":{
                "type": "object",
                "properties": {
                    "path": {
                        "type": "string",
                        "description": "Ruta del directorio"
                    }
                },
                "required": []
            }
        }
    },
    {
        "type": "function",
        "function":{
            "name": "read_file",
            "description": "Lee el contenidod de un archivo.",
            "parameters":{
                "type": "object",
                "properties": {
                    "path": {
                        "type": "string",
                        "description": "Ruta del archivo"
                    }
                },
                "required": ["path"]
            }
        }
    },
    {
        "type": "function",
        "function":{
            "name": "write_file",
            "description": "Crea o sobre escribe un archivo con el contenido indicado.",
            "parameters":{
                "type": "object",
                "properties": {
                    "path": {
                        "type": "string",
                        "description": "Ruta del archivo"
                    },
                    "content": {
                        "type": "string",
                        "description": "Contenido completo dearchivo"
                    }
                },
                "required": ["path", "content"]
            }
        }
    },
    {
        "type": "function",
        "function":{
            "name": "create_folder",
            "description": "Crea una carpeta en la ruta especificada.",
            "parameters":{
                "type": "object",
                "properties": {
                    "path": {
                        "type": "string",
                        "description": "Ruta de la carpeta a crear"
                    }
                },
                "required": ["path"]
            }
        }
    },
    {
        "type": "function",
        "function":{
            "name": "move_file_or_folder",
            "description": "Mueve un archivo o carpeta de una ubicación a otra.",
            "parameters":{
                "type": "object",
                "properties": {
                    "source": {
                        "type": "string",
                        "description": "archivo origien"
                    },
                    "destination": {
                        "type": "string",
                        "description": "carpeta de destino a donde mover el archivo, no la ruta completa con el nombre del archivo"
                    }
                },
                "required": ["source", "destination"]
            }
        }
    },
    {
        "type": "function",
        "function":{
            "name": "remove_file",
            "description": "Elimina un archivo innecesaio.",
            "parameters":{
                "type": "object",
                "properties": {
                    "file": {
                        "type": "string",
                        "description": "archivo a eliminar"
                    }
                },
                "required": ["file"]
            }
        }
    },
    {
        "type": "function",
        "function":{
            "name": "remove_folder",
            "description": "Elimina una carpeta vacía que no se esté usando. antes de eliminar la carpeta, elimina los archivos dentro de está",
            "parameters":{
                "type": "object",
                "properties": {
                    "folder": {
                        "type": "string",
                        "description": "carpeta a eliminar"
                    }
                },
                "required": ["folder"]
            }
        }
    },
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
                        "description": "llave generica para poder identificar el dato guardado, no debe repetirse, siempre debe iniciar con 'key_'"
                    },
                    "role_agent": {
                        "type": "string",
                        "description": "tú siempre deberias llamar a este parametro con el valor 'agent'",
                    },
                    "content": {
                        "type": "string",
                        "description": "dato que el agente considere importante para guardar"
                    }
                },
                "required": ["key", "content"]
            }
        }
    },
]

available_tools = {
    "list_directory": list_directory,
    "read_file": read_file,
    "write_file": write_file,
    "create_folder": create_folder,
    "move_file_or_folder": move_file_or_folder,
    "remove_file": remove_file,
    "remove_folder": remove_folder,
    "save_memory": save_memory,
}

tools_with_question = ["write_file", "remove_file", "remove_folder"]
