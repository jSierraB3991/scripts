from pathlib import Path

def remove_file(file: str) -> str:
    """Elimina un archivo innecesario de la ruta especificada."""
    file_path = Path(file)
    
    # Verificar que el archivo existe
    if not file_path.exists():
        return f"El archivo no existe: {file}"
    
    # Verificar que es efectivamente un archivo (no un directorio)
    if not file_path.is_file():
        return f"No se puede eliminar el archivo, es un directorio: {file}"
    
    try:
        file_path.unlink()
        return f"Archivo eliminado correctamente: {file}"
    except PermissionError:
        return f"No se tiene permisos para eliminar el archivo: {file}"
    except Exception as e:
        return f"Error eliminando el archivo {file}: {e}"