import os


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