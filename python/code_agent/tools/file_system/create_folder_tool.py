from pathlib import Path
import os
import logging
from typing import Optional, Union

# Configuración de logging (deshabilitado por defecto)
logger = logging.getLogger(__name__)


def create_folder(
    path: str,
    parents: bool = True,
    exist_ok: bool = False,
    mode: int = None
) -> str:
    """
    Crea una carpeta en el sistema de archivos.

    Args:
        path (str): Ruta absoluta o relativa del directorio a crear.
        parents (bool): Si es True, crea los directorios intermedios necesarios. Default: True.
        exist_ok (bool): Si es True, no lanza error si la carpeta ya existe. Default: False.
        mode (int): Modo de permisos para la carpeta creada (octal). None usa permisos por defecto.

    Returns:
        str: Mensaje indicando éxito o error.

    Raises:
        ValueError: Si el path está vacío o es inválido.
        PermissionError: Si no hay permisos para crear en la ruta.
        OSError: Para otros errores del sistema de archivos.

    Examples:
        >>> create_folder("/tmp/mi_carpeta")
        'Carpeta creada correctamente: /tmp/mi_carpeta'
        
        >>> create_folder("/path/nueva/ruta", exist_ok=True)
        'La carpeta ya existe: /path/nueva/ruta'

    """
    
    # Validación de entrada
    if not path:
        raise ValueError("La ruta del directorio no puede estar vacía")
    
    # Normalizar la ruta para evitar problemas con `.` y `..`
    normalized_path = os.path.normpath(path)
    
    # Convertir a objeto Path para operaciones
    folder_path = Path(normalized_path)

    try:
        # Verificar si el directorio ya existe
        if folder_path.exists() and folder_path.is_dir():
            if not exist_ok:
                return f"Error: La carpeta ya existe: {path}"
            
            # Obtener información sobre la carpeta existente
            stat_info = folder_path.stat()
            logger.info(f"Carapeta ya existe con permisos {oct(stat_info.st_mode)}")
            return f"La carpeta ya existe: {path} (exist_ok=True)"

        # Verificar si es un archivo, no una carpeta
        if folder_path.exists():
            return f"Error: El camino {path} apunta a un archivo, no a una carpeta"

        # Intentar crear la carpeta
        if parents:
            created = folder_path.mkdir(parents=True, mode=mode)
        else:
            # Crear solo si los padres ya existen o podemos crear toda la ruta
            try:
                created = folder_path.mkdir(mode=mode)
            except OSError as e:
                if "No such file or directory" in str(e):
                    return f"Error: Directorio padre no existe y parents=False: {path}"
                raise

        # Verificar que se creó correctamente
        if created and folder_path.exists():
            stat_info = folder_path.stat()
            permissions = oct(stat_info.st_mode)[-4:]  # Últimos 4 dígitos de modo
            logger.info(f"Carpeta creada: {path} con permisos {permissions}")
            
            return f"Carpeta creada correctamente: {path}"

        # Fallback (debería ser raro)
        if not folder_path.exists():
            return f"Error desconocido al crear carpeta: {path}"

    except PermissionError as e:
        logger.error(f"Permiso denegado para {path}: {e}")
        raise
    except OSError as e:
        error_msg = f"Error creando la carpeta {path}: {e}"
        logger.error(error_msg)
        return error_msg
    except ValueError as e:
        error_msg = f"Error de validación para {path}: {e}"
        logger.error(error_msg)
        raise
    
    # Debería no llegar aquí, pero por seguridad
    return f"Estado desconocido para: {path}"
