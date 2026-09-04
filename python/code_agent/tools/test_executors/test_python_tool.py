#!/usr/bin/env python3
"""
Executor de pruebas para Python.
Este módulo ejecuta los tests de Python."""

import os
import subprocess


def run_python_tests(path: str) -> str:
    """Ejecutar tests de Python.
    
    Args:
        path (str): Ruta del directorio donde se encuentran los tests de Python
        
    Returns:
        str: Mensaje con el resultado de la ejecución de los tests
    """
    print(f"\n🐍 Ejecutando tests de Python en: {path}")
    
    test_dir = os.path.join(path)
    if not os.path.exists(test_dir):
        return f"⚠️  Directorio de tests Python no encontrado: {test_dir}"
    
    count = subprocess.run(
        ['python3', '-m', 'pytest', test_dir, '-v'],
        capture_output=True, text=True
    )
    print(count.stdout)
    
    if count.returncode == 0:
        return "✅ Todos los tests de Python pasaron correctamente"
    else:
        return f"❌ Algunos tests de Python fallaron. Retorno: {count.returncode}"
