import sqlite3
from pathlib import Path

DB_PATH = Path(__file__).parent / "memories.db"

def get_connection():
    """Obtiene una conexión a la base de datos SQLite."""
    conn = sqlite3.connect(DB_PATH)
    return conn

def init_database():
    """Inicia la base de datos creando la tabla 'memories' si no existe."""
    conn = get_connection()
    cursor = conn.cursor()
    
    # Crear la tabla 'memories' si no existe
    cursor.execute('''
        CREATE TABLE IF NOT EXISTS memories (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            key TEXT NOT NULL,
            content TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    ''')
    
    conn.commit()
    conn.close()
