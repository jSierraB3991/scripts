from tools.file_system import create_folder_tool, list_directory_tool, move_file_or_folder_tool, read_file_tool, write_file_tool, remove_file_tool, remove_folder_tool, search_files_tool
file_system_tools = [
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
        "function": {
            "name": "search_files",
            "description": "Busca archivos dentro de un directorio. Puede bcasr por: nombre del archivo, extensión, contenido del archivo. Se puden combinar filtros",
            "parameters":{
                "type": "object",
                "properties": {
                    "path": {
                        "type": "string",
                        "description": "Directorio donde realizará la búsqueda"
                    },
                    "pattern": {
                        "type": "string",
                        "description": "Texto que debe aparecer en el nombre del archivo"
                    },
                    "content": {
                        "type": "string",
                        "description": "Texto que debe aparecer dentro del archivo"
                    },
                    "extension": {
                        "type": "string",
                        "description": "Extensión del archivo, por eje: .go o.py"
                    }
                },
                "required": ["path"]
            }
        }
    },
]

file_system_available_tools = {
    "list_directory": list_directory_tool.list_directory,
    "read_file": read_file_tool.read_file,
    "write_file": write_file_tool.write_file,
    "create_folder": create_folder_tool.create_folder,
    "move_file_or_folder": move_file_or_folder_tool.move_file_or_folder,
    "remove_file": remove_file_tool.remove_file,
    "remove_folder": remove_folder_tool.remove_folder,
    "search_files": search_files_tool.search_files,
}

