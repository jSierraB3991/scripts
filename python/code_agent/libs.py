import string
import random
import sys
import subprocess

MODEL = "cogito:8b"#"qwen3.5:9b"
MAX_LOOP_AGENT  = 15

RESET = "\033[0m"       # Volver al color normal
RED = "\033[91m"        # Rojo brillante
YELLOW = "\033[93m"     # Amarillo/Amarillo
BOLD = "\033[1m"        # Texto en negrita
PROGRAM_KEY = "python_tui_chat_llm"


KEY_NAME = "key_name_user"
KEY_STRUCTURE_PROJECT = "key_structure_proyecto"

ROL_USER = "user"
ROL_TOOL = "tool"
ROL_SYSTEM = "system"
DESCRIPT_USER_APP = "Nombre del usuario de la aplicación"

dir_no_list = ["__pycache__", ".venv", "venv", ".git", ".vscode", "node_modules", "build", "dist"]
files_no_list = [".env", ".gitignore", "memories.db"]

def random_string() ->str:
    caracteres = string.ascii_letters + string.digits 
    longitud = 15
    cadena_aleatoria = ''.join(random.choice(caracteres) for _ in range(longitud))
    return cadena_aleatoria

def clear_screen():
    subprocess.run(['cls'] if sys.platform == 'win32' else ['clear'])
