from pathlib import Path
from libs import dir_no_list, files_no_list

def list_directory(path: str = ".") -> str:
    """Lista archivos y carpetas."""
    directory = Path(path)
    if not directory.exists():
        return f"El directorio no existe: {path}"
    if not directory.is_dir():
        return f"No es un directorio: {path}"
    result = []
    for item in sorted(directory.iterdir()):
        if item.is_dir():
            if item.name not in dir_no_list:
                result.append(f"[DIR] {item.name}")
        else:
            if item.name not in files_no_list:
                result.append(f"[FILE] {item.name}")
    return "\t".join(result)