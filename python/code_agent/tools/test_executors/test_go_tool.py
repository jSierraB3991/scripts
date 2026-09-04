#!/usr/bin/env python3
"""
Executor de pruebas para Go.
Este módulo ejecuta los tests de Go."""

import os
import subprocess


def run_go_tests(path: str) -> str:
    """Ejecutar tests de Go.
    
    Args:
        path (str): Ruta del directorio donde se encuentran los tests de Go
        
    Returns:
        str: Mensaje con el resultado de la ejecución de los tests
    """
    print(f"\n🐹 Ejecutando tests de Go en: {path}")
    
    test_dir = os.path.join(path)
    if not os.path.exists(test_dir):
        return f"⚠️  Directorio de tests Go no encontrado: {test_dir}"
    
    count = subprocess.run(
        ['go', 'test', '.', '-v'],
        cwd=test_dir,
        capture_output=True, text=True
    )
    print(count.stdout)
    
    if count.returncode == 0:
        return "✅ Todos los tests de Go pasaron correctamente"
    else:
        return f"❌ Algunos tests de Go fallaron. Retorno: {count.returncode}"
