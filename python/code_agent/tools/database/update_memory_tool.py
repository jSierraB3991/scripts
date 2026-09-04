import sqlite3
from pathlib import Path
from libs import PROGRAM_KEY
from tools.database.initialize_database import initialize_database, DB_PATH

def update_memory(key: str, content: str) -> str:
    """ Actualiza el valor (contenido) en la tabla 'memories' de la base de datos SQLite. """
    if not key.lower().startswith("key_"):
        return f"la llave debe empezar con 'key_'"
    initialize_database()
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    # Insertar el contenido en la tabla 'memories'
    try:
        cursor.execute(
            "UPDATE memories SET content = ? WHERE key = ? and program_key = ?",
            (content, key, PROGRAM_KEY,)
        )
        memory_id = cursor.lastrowid
    except Exception as e:
        return f"Error actualizando la {key} en la db: error: {e}"
    
    conn.commit()
    conn.close()
    
    return f"memoria {key} actualizada correctamente"