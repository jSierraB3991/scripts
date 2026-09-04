#!/usr/bin/env python3
"""
Executor de pruebas para Go.
Este módulo ejecuta los tests de Go."""

import os
import subprocess


def run_go_tests():
    """Ejecutar tests de Go."""
    print("\n🐹 Ejecutando tests de Go...")
    
    test_dir = os.path.join(os.path.dirname(__file__), '..', 'tests', 'go_tests')
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
