from pathlib import Path

def move_file_or_folder(source: str, destination: str) -> str:
    """Mueve un archivo o carpeta de una ubicación a otra."""
    source_path = Path(source).resolve()
    destination_path = Path(destination).resolve()
    
    # Verificar que el origen existe
    if not source_path.exists():
        return f"El origen no existe: {source}"
    
    # Verificar que el destino es un directorio
    if not destination_path.is_dir():
        try:
            destination_path.mkdir(parents=True, exist_ok=True)
        except Exception as e:
            return f"Error creando la carpeta de destino: {destination}: {e}"
    
    # Verificar que el origen es archivo o carpeta (no ambos)
    if source_path.is_file() and not source_path.is_dir():
        try:
            source_path.rename(destination_path.joinpath(source_path.name))
            return f"Archivo movido correctamente desde '{source}' a '{destination}'"
        except Exception as e:
            return f"Error moviendo el archivo {source}: {e}"
    elif source_path.is_dir() and not source_path.is_file():
        try:
            source_path.rename(destination_path.joinpath(source_path.name))
            return f"Carpeta movida correctamente desde '{source}' a '{destination}'"
        except Exception as e:
            return f"Error moviendo la carpeta {source}: {e}"
    else:
        return "No es posible mover tanto archivo como carpeta al mismo tiempo."