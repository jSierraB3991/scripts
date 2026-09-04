import sqlite3
from pathlib import Path
from libs import PROGRAM_KEY

# Importar funciones de db.py
DB_PATH = Path(__file__).parent.parent / "memories.db"


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
            role_agent TEXT NOT NULL,
            program_key TEXT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    """)
    
    conn.commit()
    conn.close()


def save_memory(key: str, role_agent: str, content: str) -> str:
    """ Guarda un valor (contenido) en la tabla 'memories' de la base de datos SQLite. """
    initialize_database()
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    # Insertar el contenido en la tabla 'memories'
    try:
        cursor.execute(
            "INSERT INTO memories (key, content, role_agent, program_key) VALUES (?,?,?)",
            (key, content,role_agent, PROGRAM_KEY)
        )
        memory_id = cursor.lastrowid
    except sqlite3.IntegrityError:
        return f"La clave '{key}' ya existe en la base de datos."
    except Exception as e:
        return f"Error guardado {key} en la db: error: {e}"
    
    conn.commit()
    conn.close()
    
    return f"memoria {key} guardada con el id: {memory_id}"


def get_all_memories() -> list["Memory"]:
    """ Obtiene todas las memorias guardadas en la base de datos. """
    initialize_database()
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute("SELECT id, key, content, role_agent, created_at FROM memories WHERE program_key = ?", (PROGRAM_KEY,))
    rows = cursor.fetchall()

    # Convertir a lista de objetos Memory
    memories = [Memory(*row) for row in rows]
    
    conn.close()
    return memories


def get_memory_by_id(memory_id: int) -> Memory | None:
    """ Obtiene una memoria por su ID. """
    initialize_database()
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute(
        "SELECT id, key, content, role_agent, created_at FROM memories WHERE id = ? and program_key = ?",
        (memory_id,PROGRAM_KEY,)
    )
    row = cursor.fetchone()
    
    if row:
        memory = Memory(**row)
    else:
        memory = None
    
    conn.close()
    return memory


def get_memory_by_key(key: str) -> Memory | None:
    """ Obtiene una memoria por su clave (key). """
    initialize_database()
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute(
        "SELECT id, key, content, role_agent, created_at FROM memories WHERE key = ? and program_key = ?",
        (key,PROGRAM_KEY,)
    )
    row = cursor.fetchone()
    
    if row:
        memory = Memory(id=row[0],key=row[1],content=row[2], role_agent=row[3], created_at=row[4])
    else:
        memory = None
    
    conn.close()
    return memory


class Memory:
    """Clase que representa una memoria guardada en la base de datos."""
    
    def __init__(self, id: int, key: str, content: str, role_agent: str, created_at: str):
        self.id = id
        self.key = key
        self.content = content
        self.created_at = created_at
        self.role_agent = role_agent
    
    def to_dict(self) -> dict:
        """Convierte el objeto Memory a un diccionario."""
        return {
            "id": self.id,
            "key": self.key,
            "content": self.content,
            "created_at": self.created_at,
            "role_agent": self.role_agent,
        }