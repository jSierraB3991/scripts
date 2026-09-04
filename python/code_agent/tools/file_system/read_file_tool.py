from pathlib import Path

def read_file(path: str) -> str:
    """Lee el contenido de un archivo."""
    file = Path(path)
    if not file.exists():
        return f"El archivo no existe: {path}"
    if not file.is_file():
        return f"No es un archivo: {path}"

    try:
        return file.read_text(encoding="utf-8")
    except Exception as e:
        return f"Error leyendo el archivo: {e}"