from pathlib import Path

def create_folder(path: str = ".") -> str:
    """Crea una carpeta."""
    folder = Path(path)
    try:
        if folder.exists():
            return f"La carpeta ya existe: {path}"
        folder.mkdir(parents=True, exist_ok=True)
        return f"Carpeta creada correctamente: {path}"
    except Exception as e:
        return f"Error creando la carpeta {path}: {e}"