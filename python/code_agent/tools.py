from pathlib import Path

def list_directory(path: str = ".") -> str:
    """Lista archivos y carpetas."""
    directory = Path(path)
    if not directory.exists():
        return f"El archivo no existe: {path}"
    if not directory.is_dir():
        return f"No es un directory: {path}"
    result = []
    for item in sorted(directory.iterdir()):
        if item.is_dir():
            result.append(f"[DIR] {item.name}")
        else:
            result.append(f"[FILE] {item.name}")
    return "\n".join(result)


def read_file(path: str) -> str:
    """Lee el contenido de un archivo."""
    file = Path(path)
    if not file.exists():
        return f"El arhivo no existe: {path}"
    if not file.is_file():
        return f"No es un archivo: {path}"

    try:
        return file.read_text(encoding="utf-8")
    except Exception as e:
        return f"Error leyendo el archivo: {e}"

def write_file(path: str, content: str) -> str:
    "Escribe el contenido en en archivo."
    file = Path(path)
    try:
        file.parent.mkdir(parents=True, exist_ok=True)
        file.write_text(content, encoding="utf-8")
        return f"Archivo escrito correctamente: {path}"
    except Exception as e:
        return f"Error escribiendo en el archivo {path}"


