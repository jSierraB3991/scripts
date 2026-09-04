#!/usr/bin/env python3
"""
Executor de pruebas para Python.
Este módulo ejecuta los tests de Python."""

import os
import subprocess


def run_python_tests():
    """Ejecutar tests de Python."""
    print("\n🐍 Ejecutando tests de Python...")
    
    # Buscar archivos .py en python_tests/
    test_dir = os.path.join(os.path.dirname(__file__), '..', 'tests', 'python_tests')
    if not os.path.exists(test_dir):
        print("⚠️  Directorio de tests Python no encontrado.")
        return
    
    count = subprocess.run(
        ['python3', '-m', 'pytest', test_dir, '-v'],
        capture_output=True, text=True
    )
    print(count.stdout)
    
    return count.returncode == 0
