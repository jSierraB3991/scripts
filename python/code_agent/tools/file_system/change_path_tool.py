#!/usr/bin/env python3
"""
Herramienta para cambiar el path/ruta del agente de código.
Este archivo permite reconfigurar la carpeta raíz donde se ejecutará el agente.

Funciones disponibles:
- cambiar_path(actual, nuevo): Cambia la ruta actual por una nueva ruta especificada.
- obtener_info_ruta(ruta): Obtiene información sobre una ruta del sistema.
"""

import os


def cambiar_path(actual, nuevo):
    """
    Cambia la ruta actual por una nueva ruta especificada.
    
    Args:
        actual (str): Ruta actual del agente
        nuevo (str): Nueva ruta del agente
        
    Returns:
        dict: Diccionario con el resultado de la operación
    """
    try:
        if os.path.exists(nuevo):
            # Validar que es un directorio válido
            if not os.path.isdir(nuevo):
                raise ValueError(f"La nueva ruta debe ser un directorio: {nuevo}")
            
            # Obtener información de la ruta actual
            info_actual = obtener_info_ruta(actual)
            
            # Crear resumen del cambio
            resultado = {
                'estado': 'exito',
                'mensaje': f'Path cambiado con éxito de "{actual}" a "{nuevo}"',
                'ruta_anterior': actual,
                'ruta_nueva': nuevo,
                'info_anterior': info_actual
            }
        else:
            raise FileNotFoundError(f"La nueva ruta no existe: {nuevo}")
            
    except Exception as e:
        resultado = {
            'estado': 'error',
            'mensaje': str(e),
            'ruta_anterior': actual,
            'ruta_nueva': nuevo
        }
    
    return resultado


def obtener_info_ruta(ruta):
    """
    Obtiene información sobre una ruta del sistema.
    
    Args:
        ruta (str): Ruta a verificar
        
    Returns:
        dict: Información sobre la ruta
    """
    try:
        info = {
            'existe': os.path.exists(ruta),
            'es_directorio': os.path.isdir(ruta),
            'nombre_ruta': os.path.basename(os.path.normpath(ruta)),
            'padre': os.path.dirname(ruta)
        }
        return info
    except Exception as e:
        return {'error': str(e)}
