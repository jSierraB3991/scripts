from tools import *
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
]

available_tools = {
    "list_directory": list_directory,
    "read_file": read_file,
    "write_file": write_file,
    "create_folder": create_folder,
}
