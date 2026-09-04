import sqlite3
from pathlib import Path
from libs import PROGRAM_KEY
from models.memory import Memory
from tools.database.initialize_database import initialize_database, DB_PATH

def get_memory_by_key(key: str) -> Memory | None:
    """ Obtiene una memoria por su clave (key). """
    initialize_database()
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute(
        "SELECT id, key, content, role_agent, description, created_at FROM memories WHERE key = ? and program_key = ?",
        (key,PROGRAM_KEY,)
    )
    row = cursor.fetchone()
    
    if row:
        memory = Memory(id=row[0],key=row[1],content=row[2], role_agent=row[3], description=row[4],created_at=row[5])
    else:
        memory = None
    
    conn.close()
    return memory


