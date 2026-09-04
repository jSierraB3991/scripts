import sqlite3
from pathlib import Path
from libs import PROGRAM_KEY

# Importar funciones de db.py
DB_PATH = Path(__file__).parent.parent.parent / "memories.db"


def initialize_database():
    """Inicializa la base de datos si no existe."""
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    # Crear tabla memories si no existe
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS memories (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            key TEXT UNIQUE NOT NULL,
            content TEXT NOT NULL,
            description TEXT  NOT NULL,
            role_agent TEXT NOT NULL,
            program_key TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    """)
    
    conn.commit()
    conn.close()