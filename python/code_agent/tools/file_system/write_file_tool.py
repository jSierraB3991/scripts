from pathlib import Path

def write_file(path: str, content: str) -> str:
    """Escribe el contenido en un archivo."""
    file = Path(path)
    try:
        file.parent.mkdir(parents=True, exist_ok=True)
        file.write_text(content, encoding="utf-8")
        return f"Archivo escrito correctamente: {path}"
    except Exception as e:
        return f"Error escribiendo en el archivo {path}"