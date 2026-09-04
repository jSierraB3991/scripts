import sqlite3
from pathlib import Path
from libs import PROGRAM_KEY
from models.memory import Memory
from tools.database.initialize_database import initialize_database, DB_PATH


def get_all_memories() -> list["Memory"]:
    """ Obtiene todas las memorias guardadas en la base de datos. """
    initialize_database()
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute("SELECT id, key, content, role_agent, description, created_at FROM memories WHERE program_key = ?", (PROGRAM_KEY,))
    rows = cursor.fetchall()

    # Convertir a lista de objetos Memory
    memories = [Memory(*row) for row in rows]
    
    conn.close()
    return memories
