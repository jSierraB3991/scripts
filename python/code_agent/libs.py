import string
import random
import sys
import os
import subprocess

MODEL = "gemma4:e4b"#"qwen3.5:9b"

RESET = "\033[0m"       # Volver al color normal
RED = "\033[91m"        # Rojo brillante
YELLOW = "\033[93m"     # Amarillo/Amarillo
BOLD = "\033[1m"        # Texto en negrita


KEY_NAME = "key_name_user"
KEY_SYSTEM_PROMPT = "maximun_prompt"

ROL_USER = "user"
ROL_TOOL = "tool"
ROL_SYSTEM = "system"

def random_string() ->str:
    caracteres = string.ascii_letters + string.digits 
    longitud = 15
    cadena_aleatoria = ''.join(random.choice(caracteres) for _ in range(longitud))
    return cadena_aleatoria

def clear_screen():
    subprocess.run(['cls'] if sys.platform == 'win32' else ['clear'])