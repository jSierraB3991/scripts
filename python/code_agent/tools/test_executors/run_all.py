#!/usr/bin/env python3
"""
Executor de pruebas para todos los lenguajes.
Este script coordinará la ejecución de tests en Python y Go.
"""

import os
import subprocess


def run_python_tests():
    """Ejecutar tests de Python."""
    print("\n🐍 Ejecutando tests de Python...")
    
    # Buscar archivos .py en python_tests/
    test_dir = os.path.join(os.path.dirname(__file__), '..', '..', 'tests', 'python_tests')
    if not os.path.exists(test_dir):
        print("⚠️  Directorio de tests Python no encontrado.")
        return
    
    count = subprocess.run(
        ['python3', '-m', 'pytest', test_dir, '-v'],
        capture_output=True, text=True
    )
    print(count.stdout)
    
    return count.returncode == 0


def run_go_tests():
    """Ejecutar tests de Go."""
    print("\n🐹 Ejecutando tests de Go...")
    
    test_dir = os.path.join(os.path.dirname(__file__), '..', '..', 'tests', 'go_tests')
    if not os.path.exists(test_dir):
        print("⚠️  Directorio de tests Go no encontrado.")
        return
    
    count = subprocess.run(
        ['go', 'test', '.', '-v'],
        cwd=test_dir,
        capture_output=True, text=True
    )
    print(count.stdout)
    
    return count.returncode == 0


# La lógica de ejecución principal ha sido eliminada.
# Este archivo ahora solo contiene funciones utilitarias para ejecutar tests.
