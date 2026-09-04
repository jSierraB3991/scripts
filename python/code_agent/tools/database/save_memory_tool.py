import sqlite3
from pathlib import Path
from libs import PROGRAM_KEY
from tools.database.initialize_database import initialize_database, DB_PATH



def save_memory(key: str, role_agent: str, content: str,description:str = "") -> str:
    """ Guarda un valor (contenido) en la tabla 'memories' de la base de datos SQLite. """
    initialize_database()
    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()

    if description == "":
        description = key
    
    # Insertar el contenido en la tabla 'memories'
    try:
        cursor.execute(
            "INSERT INTO memories (key, content, role_agent, program_key,description) VALUES (?,?,?,?,?)",
            "INSERT INTO memories (key, content, role_agent, program_key,description) VALUES (?,?,?,?,?)",
            (key, content,role_agent, PROGRAM_KEY,description,)
        )
        memory_id = cursor.lastrowid
    except sqlite3.IntegrityError:
        return f"Error: La clave '{key}' ya existe en la base de datos, si recuerdas que tienes una tool para actualizar?."
    except Exception as e:
        return f"Error guardado {key} en la db: error: {e}"
    
    conn.commit()
    conn.close()
    
    return f"memoria {key} guardada con el id: {memory_id}"