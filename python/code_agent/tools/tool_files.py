from pathlib import Path
import os

dir_no_list = ["__pycache__", ".venv", "venv", ".git"]
files_no_list = [".env", ".gitignore", "**.db"]
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
    return "\n".join(result)


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


def write_file(path: str, content: str) -> str:
    """Escribe el contenido en un archivo."""
    file = Path(path)
    try:
        file.parent.mkdir(parents=True, exist_ok=True)
        file.write_text(content, encoding="utf-8")
        return f"Archivo escrito correctamente: {path}"
    except Exception as e:
        return f"Error escribiendo en el archivo {path}"


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

def remove_folder(folder) -> str:
    """ Función auxiliar para eliminar solo una carpeta vacía (sin contenido). """
    
    # Validar ruta
    if not folder or not isinstance(folder, str) or len(folder.strip()) == 0:
        return f"Error: La ruta debe ser una cadena no vacía. Recibido: '{folder}'"
    
    folder = os.path.normpath(folder)
    
    if not os.path.exists(folder):
        return f"Advertencia: La carpeta no existe: {folder}"
    
    try:
        # Solo podemos eliminar directorios vacíos con os.rmdir()
        os.rmdir(folder)
        return f"Éxito: Carpeta vacía eliminada correctamente: {folder}"
        
    except OSError as e:
        if e.errno == 39 or "Directory not empty" in str(e):
            return f"Error: La carpeta {folder} no está vacía: {e}"
        else:
            return f"Error al eliminar la carpeta vacía {folder}: {type(e).__name__} - {e}"