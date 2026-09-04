import sqlite3
from pathlib import Path

# Importar funciones de db.py
DB_PATH = Path(__file__).parent.parent / "memories.db"

class Memory:
    def __init__(self, id, key, content, created):
        self.id = id
        self.key= key
        self.content = content
        self.created = created


def save_memory(key:str, content: str) -> int:
    """
    Guarda un valor (contenido) en la tabla 'memories' de la base de datos SQLite.
    
    Parámetros:
        content (str): El contenido/texto que se quiere guardar en la tabla memories.
    
    Retorna:
        int: El ID de la memoria guardada.
    """
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    # Insertar el contenido en la tabla 'memories'
    cursor.execute(
        "INSERT INTO memories (key, content) VALUES (?,?)",
        (key, content,)
    )
    
    # Obtener el ID de la última inserción
    memory_id = cursor.lastrowid
    
    conn.commit()
    conn.close()
    
    return memory_id


def get_all_memories() -> list[Memory]:
    """
    Obtiene todas las memorias guardadas en la base de datos.
    
    Retorna:
        list: Una lista de diccionarios con los registros de la tabla 'memories'.
    """
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute("SELECT id, key, content, created_at FROM memories")
    rows = cursor.fetchall()

    memories: list[Memory]
    # Convertir a lista de diccionarios
    for row in rows:
        pass
    
    conn.close()
    return memories


def get_memory_by_id(memory_id: int) -> dict | None:
    """
    Obtiene una memoria por su ID.
    
    Parámetros:
        memory_id (int): El ID de la memoria a buscar.
    
    Retorna:
        dict: Un diccionario con la memoria encontrada o None si no existe.
    """
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute(
        "SELECT id, key, content, created_at FROM memories WHERE id = ?",
        (memory_id,)
    )
    row = cursor.fetchone()
    
    if row:
        memory = {
            "id": row[0],
            "key": row[1],
            "content": row[2],
            "created_at": row[3]
        }
    else:
        memory = None
    
    conn.close()
    return memory


def get_memory_by_key(key: str) -> dict | None:
    """
    Obtiene una memoria por su ID.
    
    Parámetros:
        memory_id (int): El ID de la memoria a buscar.
    
    Retorna:
        dict: Un diccionario con la memoria encontrada o None si no existe.
    """
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute(
        "SELECT id, key, content, created_at FROM memories WHERE key = ?",
        (key,)
    )
    row = cursor.fetchone()
    
    if row:
        memory = {
            "id": row[0],
            "key": row[1],
            "content": row[2],
            "created_at": row[3]
        }
    else:
        memory = None
    
    conn.close()
    return memory

