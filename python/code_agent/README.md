# Code Agent

Agente de código local desarrollado en Python utilizando **Ollama**, con soporte para herramientas (_tools_), memoria persistente y ejecución interactiva desde terminal.

El agente está diseñado para ayudar a desarrollar y modificar proyectos utilizando herramientas controladas para interactuar con el sistema de archivos, memoria del proyecto y ejecución de pruebas.

## Características

- 🤖 Ejecución de modelos locales mediante Ollama.
- 🧠 Memoria persistente almacenada en base de datos.
- 📁 Lectura, creación, modificación, movimiento y eliminación de archivos.
- 🔎 Búsqueda de archivos por nombre, extensión o contenido.
- 🧪 Ejecución de pruebas de Go y Python.
- 🔐 Confirmación del usuario para operaciones sensibles.
- 🔄 Reinicio automático del contexto cuando se alcanza el límite de iteraciones.
- 💻 Interfaz interactiva desde terminal.
- 📋 Barra de progreso durante las consultas al modelo.
- 🏗️ Memoria dedicada para mantener la estructura del proyecto.

---

## Arquitectura

El agente está dividido principalmente en tres grupos de herramientas:

```text
Agent
│
├── Database Tools
│   ├── save_memory
│   ├── update_memory
│   └── get_memory_by_key_tool
│
├── File System Tools
│   ├── list_directory
│   ├── read_file
│   ├── write_file
│   ├── create_folder
│   ├── move_file_or_folder
│   ├── remove_file
│   ├── remove_folder
│   └── search_files
│
└── Test Tools
    ├── run_go_tests
    └── run_python_tests
```

Las definiciones que Ollama recibe y las funciones Python que realmente ejecutan las operaciones están separadas.

```text
Tool Definition
      │
      ▼
my_tools
      │
      ▼
Ollama decide utilizar una tool
      │
      ▼
available_tools
      │
      ▼
Función Python
      │
      ▼
Resultado
      │
      ▼
Agent
```

---

## Estructura del proyecto

Una estructura esperada para el proyecto es:

```text
.
├── agent.py
├── libs.py
├── progress_bar.py
│
└── tools/
    ├── tools_descript.py
    │
    ├── database/
    │   ├── database_tool_descript.py
    │   ├── save_memory_tool.py
    │   ├── update_memory_tool.py
    │   ├── get_memory_by_key_tool.py
    │   └── get_all_memories.py
    │
    ├── file_system/
    │   ├── file_system_tool_descript.py
    │   ├── create_folder_tool.py
    │   ├── list_directory_tool.py
    │   ├── move_file_or_folder_tool.py
    │   ├── read_file_tool.py
    │   ├── write_file_tool.py
    │   ├── remove_file_tool.py
    │   ├── remove_folder_tool.py
    │   └── search_files_tool.py
    │
    └── test_executors/
        ├── test_tool_descript.py
        ├── test_go_tool.py
        └── test_python_tool.py
```

La estructura real puede variar dependiendo de cómo esté organizado el proyecto.

---

# Agente

La clase principal es:

```python
class Agent:
```

Su responsabilidad es administrar:

- contexto del modelo;
- memoria;
- ejecución de herramientas;
- confirmación de operaciones;
- interacción con el usuario;
- límite de iteraciones.

---

## Inicialización

Al crear el agente:

```python
agent = Agent()
```

se inicializa el contador de iteraciones y posteriormente la memoria.

```python
def __init__(self):
    self.loop = 0
    self.init_memory()
```

`init_memory()` crea el mensaje inicial del sistema y carga la información persistente necesaria.

---

# System Prompt

El agente recibe instrucciones iniciales que establecen su comportamiento.

Entre ellas:

- trabajar como agente de código;
- utilizar únicamente las herramientas disponibles;
- leer archivos antes de modificarlos cuando sea necesario;
- no inventar contenido de archivos que no haya leído;
- trabajar dentro de la raíz del proyecto;
- mantener conocimiento de la estructura del proyecto;
- no crear `__init__.py` en proyectos Python;
- utilizar la memoria cuando sea necesario.

Esto permite que las restricciones principales del agente estén definidas desde el inicio de la conversación.

---

# Memoria

El agente utiliza una base de datos para almacenar información persistente.

Existen tres operaciones principales.

## `save_memory`

Guarda una nueva memoria.

Parámetros:

```text
key
role_agent
content
description
```

Las claves deben comenzar con:

```text
key_
```

Ejemplo conceptual:

```json
{
    "key": "key_project_structure",
    "role_agent": "agent",
    "content": "...",
    "description": "Estructura del proyecto"
}
```

---

## `update_memory`

Actualiza una memoria existente utilizando su `key`.

```text
key
content
```

Esta herramienta permite modificar información previamente almacenada.

---

## `get_memory_by_key_tool`

Obtiene una memoria específica utilizando su clave.

```text
key
```

Es útil cuando el agente necesita recuperar rápidamente información concreta.

---

# Estructura del proyecto en memoria

El agente puede almacenar la estructura del proyecto utilizando una clave definida en:

```python
libs.KEY_STRUCTURE_PROJECT
```

Esto permite que el agente tenga una referencia persistente de dónde se encuentran las carpetas y archivos importantes.

Si la estructura cambia, el agente puede actualizar la memoria correspondiente mediante:

```text
update_memory
```

---

# Herramientas del sistema de archivos

Las herramientas de archivos permiten al agente interactuar con el proyecto sin darle acceso directo a una shell.

## `list_directory`

Lista archivos y carpetas de un directorio.

```text
path
```

---

## `read_file`

Lee el contenido de un archivo.

```text
path
```

Esta herramienta es especialmente importante antes de modificar archivos existentes.

---

## `write_file`

Crea o sobrescribe un archivo.

```text
path
content
```

Esta operación requiere confirmación del usuario.

---

## `create_folder`

Crea una carpeta.

```text
path
```

---

## `move_file_or_folder`

Mueve un archivo o carpeta.

```text
source
destination
```

---

## `remove_file`

Elimina un archivo.

```text
file
```

Esta operación requiere confirmación.

---

## `remove_folder`

Elimina una carpeta.

```text
folder
```

Esta operación requiere confirmación.

---

## `search_files`

Busca archivos dentro de un directorio.

Puede utilizar diferentes filtros:

```text
path
pattern
content
extension
```

Los filtros pueden combinarse.

Ejemplos conceptuales:

```text
Buscar archivos cuyo nombre contenga "user"
```

```text
Buscar archivos .py
```

```text
Buscar archivos que contengan "get_memory"
```

---

# Herramientas de pruebas

El agente puede ejecutar pruebas de diferentes lenguajes.

## Go

```text
run_go_tests
```

Recibe:

```text
path
```

y ejecuta los tests correspondientes al directorio indicado.

## Python

```text
run_python_tests
```

Recibe:

```text
path
```

y ejecuta los tests correspondientes al directorio indicado.

---

# Confirmación de herramientas

No todas las herramientas requieren autorización.

Las herramientas consideradas sensibles están definidas en:

```python
tools_with_question = [
    "write_file",
    "remove_file",
    "remove_folder",
    "update_memory"
]
```

Cuando el agente intenta utilizar una de ellas, se solicita confirmación:

```text
voy a usar la herramienta: write_file ...
está de acuerdo? Y/n:
```

Si el usuario responde:

```text
y
```

la herramienta se ejecuta.

Si responde cualquier otra cosa, se solicita una razón y dicha información se incorpora al contexto del agente.

Esto permite que el usuario mantenga el control sobre operaciones destructivas o modificaciones importantes.

---

# Flujo de ejecución

Cuando el usuario realiza una consulta:

```text
Usuario
   │
   ▼
run_agent()
   │
   ▼
Se agrega el mensaje al contexto
   │
   ▼
Ollama
   │
   ├── Respuesta normal
   │      │
   │      ▼
   │   Mostrar respuesta
   │
   └── Tool call
          │
          ▼
       run_tools()
          │
          ▼
    ¿Requiere confirmación?
       │          │
      Sí          No
       │          │
       ▼          ▼
   Preguntar    Ejecutar
       │          │
       └────┬─────┘
            ▼
      Resultado de tool
            │
            ▼
      Se agrega al contexto
            │
            ▼
         Ollama
```

El ciclo continúa hasta que el modelo genera una respuesta final sin solicitar otra herramienta.

---

# Control del contexto

El agente mantiene sus mensajes en:

```python
self.messages
```

Cuando se alcanza el límite establecido por:

```python
libs.MAX_LOOP_AGENT
```

se reinicia el contexto utilizando una copia previamente guardada:

```python
self.backup
```

El objetivo es evitar que una conversación demasiado larga consuma indefinidamente el contexto del modelo.

---

# Interfaz de terminal

El agente proporciona una interfaz interactiva mediante:

```python
prepare(name_user)
```

Ejemplo:

```text
Escribe 'exit' para salir.

Usuario - modelo >
```

Comandos disponibles:

| Comando   | Acción                     |
| --------- | -------------------------- |
| `exit`    | Salir                      |
| `quit`    | Salir                      |
| `clear`   | Limpiar pantalla           |
| `cls`     | Limpiar pantalla           |
| `refresh` | Reiniciar memoria/contexto |

Cualquier otro texto se envía al agente.

---

# Ollama

El modelo utilizado se configura mediante:

```python
libs.MODEL
```

La comunicación se realiza mediante:

```python
ollama.chat(
    model=libs.MODEL,
    messages=self.messages,
    tools=my_tools
)
```

Esto permite que el modelo reciba tanto el historial de conversación como las definiciones de las herramientas disponibles.

---

# Separación entre definición y ejecución

Una característica importante del diseño es que las herramientas tienen dos partes.

### Definición

Es lo que recibe Ollama:

```python
{
    "type": "function",
    "function": {
        "name": "read_file",
        "description": "...",
        "parameters": {...}
    }
}
```

### Implementación

Es la función Python que realmente ejecuta la operación:

```python
file_system_available_tools = {
    "read_file": read_file_tool.read_file,
}
```

Esto permite que el modelo únicamente conozca las operaciones que puede solicitar, mientras que Python controla cómo se ejecutan realmente.

---

# Filosofía del agente

El proyecto sigue una arquitectura donde:

```text
LLM
 │
 │ decide qué necesita hacer
 ▼
Tool
 │
 │ ejecuta una operación controlada
 ▼
Resultado
 │
 ▼
LLM
```

El modelo **no recibe acceso directo a una shell**. En lugar de eso, las operaciones disponibles se exponen explícitamente como herramientas Python.

Esto permite ampliar las capacidades del agente agregando nuevas tools sin modificar necesariamente el núcleo de `Agent`.

---

# Agregar una nueva herramienta

Una nueva categoría de herramientas debería tener:

```text
tools/
└── nueva_categoria/
    ├── nueva_categoria_tool_descript.py
    └── ...
```

La definición debe agregarse a la lista de tools:

```python
nueva_categoria_tools = [...]
```

y su implementación al diccionario:

```python
nueva_categoria_available_tools = {
    "nombre_tool": funcion_python
}
```

Después se incorporan a:

```python
my_tools
```

y:

```python
available_tools
```

respectivamente.

Si la operación requiere autorización, también debe agregarse a:

```python
tools_with_question
```

---

# Dependencias principales

El proyecto utiliza principalmente:

- Python
- Ollama
- `ollama` Python SDK
- sistema de archivos local
- base de datos utilizada por el sistema de memoria

Las dependencias adicionales dependen de las implementaciones concretas de las tools.

---

# Objetivo

El objetivo del proyecto es construir un **agente de código local extensible**, capaz de utilizar un modelo ejecutado mediante Ollama para razonar sobre un proyecto y realizar acciones mediante herramientas controladas.

La arquitectura permite agregar progresivamente nuevas capacidades como:

```text
       ┌───────────────┐
       │      LLM      │
       └───────┬───────┘
               │
       ┌───────▼───────┐
       │     Agent     │
       └───────┬───────┘
               │
    ┌──────────┼──────────┐
    │          │          │
    ▼          ▼          ▼
 Memoria    Archivos    Tests
    │          │          │
    ▼          ▼          ▼
 Database   Proyecto    Go/Python
```

La intención es mantener el núcleo del agente pequeño y delegar las capacidades específicas en herramientas independientes.
