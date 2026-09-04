from pathlib import Path

def search_files(path: str = ".", pattern: str = "", content: str ="", extension: str = "") -> str:
    """Busca archivos por nombre/extensión y/o contenido"""
    path_root = Path(path)
    if not path_root.exists():
        return f"La ruta no existe: {path}"
    if not path_root.is_dir():
        return f"La ruta no es un directorio: {path}"
    result = []
    try:
        for file in path_root.rglob("*"):
            if not file.is_file():
                continue

            #ignorar directorios comunes de configuración
            if any(part in dir_no_list for part in file.parts):
                continue

            #Buscar por nombre
            if pattern and pattern.lower() not in file.name.lower():
                continue

            #Buscar por extensión
            if extension and file.suffix.lower() != extension.lower():
                continue

            #Buscar dentro del archivo
            if content:
                try:
                    text = file.read_text(encoding="utf-8",errors="ignore")
                except Exception:
                    continue

                if content.lower() not in text.lower():
                    continue
            result.append(str(file))

    except PermissionError as e:
        return f"Error: no tien permisos para acceder a algunas rutas. {e}"
    except Exception as e:
        return f"Error search files {e}"
    return "\t".join(result)
    