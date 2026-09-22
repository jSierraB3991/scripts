# Ollama TUI

TUI para interactuar con modelos locales de **Ollama** desde la terminal, construida con **Python + Textual**.

Incluye selección dinámica de modelos, streaming de respuestas y un sistema de memoria persistente capaz de resumir conversaciones antiguas.

## Características

- 🖥️ Interfaz de terminal construida con **Textual**.
- 🤖 Integración con **Ollama**.
- 🔄 Detección automática de los modelos instalados.
- 🎛️ Cambio de modelo desde la interfaz.
- 💬 Conversación interactiva.
- 🧠 Memoria persistente en `memory.json`.
- 📝 Resumen automático de conversaciones antiguas.
- ⚡ Respuestas mediante streaming.
- 🧵 Consultas a Ollama ejecutadas en workers para no bloquear la interfaz.
- ⌨️ `Ctrl+C` para salir.

---

## Requisitos

- Python 3.10+
- Ollama instalado y ejecutándose.
- Al menos un modelo disponible en Ollama.

Instalar dependencias:

```bash
pip install textual ollama
```

---

## Estructura

```text
.
├── main.py
├── simple_memory.py
├── memory.json
└── README.md
```

### `main.py`

Contiene la interfaz TUI y la comunicación principal con Ollama.

### `simple_memory.py`

Implementa el sistema de memoria de la aplicación.

### `memory.json`

Almacena de forma persistente:

- El resumen de conversaciones anteriores.
- Los mensajes recientes.

El archivo se genera automáticamente cuando se guardan mensajes.

---

# Memoria

El proyecto utiliza `SimpleMemory` para mantener contexto entre conversaciones.

```python
memory = SimpleMemory()
```

La memoria tiene dos componentes:

```text
┌──────────────────────────────┐
│        SimpleMemory          │
├──────────────────────────────┤
│ Summary                      │
│ Resumen consolidado          │
├──────────────────────────────┤
│ Recent messages              │
│ Mensajes recientes           │
└──────────────────────────────┘
```

Esto permite mantener información importante sin enviar indefinidamente todo el historial al modelo.

## Memoria reciente

Por defecto se mantienen como máximo **12 mensajes**:

```python
SimpleMemory(max_msg=12)
```

Cada mensaje se almacena con el formato utilizado por Ollama:

```python
{
    "role": "user",
    "content": "Hola"
}
```

o:

```python
{
    "role": "assistant",
    "content": "¡Hola!"
}
```

## Resumen automático

Cuando la cantidad de mensajes supera `max_msg`, se toman los primeros `num_summarize` mensajes y se envían a Ollama para generar un resumen.

Configuración predeterminada:

```python
SimpleMemory(
    max_msg=12,
    num_summarize=10
)
```

El flujo es:

```text
Nuevos mensajes
      │
      ▼
¿Supera max_msg?
      │
     Sí
      │
      ▼
Tomar los primeros
num_summarize mensajes
      │
      ▼
Ollama genera resumen
      │
      ▼
Guardar resumen
      │
      ▼
Conservar mensajes recientes
```

El resumen también tiene en cuenta el resumen anterior, permitiendo consolidar progresivamente la información.

## Contexto enviado al modelo

Si existe un resumen, se convierte en un mensaje `system`:

```python
{
    "role": "system",
    "content": "Memorias de conversación: ..."
}
```

Después se añaden los mensajes recientes.

Por tanto, Ollama recibe algo equivalente a:

```text
┌──────────────────────────────┐
│ System                       │
│ Memorias de conversación     │
├──────────────────────────────┤
│ User                         │
│ Mensaje reciente             │
├──────────────────────────────┤
│ Assistant                    │
│ Respuesta reciente           │
├──────────────────────────────┤
│ User                         │
│ Mensaje actual               │
└──────────────────────────────┘
```

---

# Memoria persistente

La información se guarda en:

```text
memory.json
```

Ejemplo:

```json
{
    "summary": "El usuario está desarrollando una aplicación en Python...",
    "memory": [
        {
            "role": "user",
            "content": "Quiero agregar herramientas"
        },
        {
            "role": "assistant",
            "content": "Puedes implementar..."
        }
    ]
}
```

Al iniciar la aplicación, `SimpleMemory` intenta cargar automáticamente este archivo.

Si no existe, la aplicación puede comenzar con una memoria vacía.

---

# Modelo utilizado para resumir

El modelo utilizado para generar los resúmenes está definido por:

```python
DEFAULT_MODEL = "lfm2.5:8b"
```

El resumen se genera mediante Ollama independientemente del modelo seleccionado actualmente en la interfaz.

Esto significa que:

```text
Modelo seleccionado
       │
       └──► Responde al usuario

DEFAULT_MODEL
       │
       └──► Genera los resúmenes de memoria
```

---

# Uso

Inicia Ollama:

```bash
ollama serve
```

Comprueba los modelos disponibles:

```bash
ollama list
```

Por ejemplo:

```bash
ollama pull lfm2.5:8b
```

Ejecuta la aplicación:

```bash
python main.py
```

---

# Interfaz

```text
┌─────────────────────────────────────────────────┐
│                    Ollama TUI                   │
├─────────────────────────────────────────────────┤
│                                                 │
│ Tú: Hola                                        │
│                                                 │
│ lfm2.5:8b: ¡Hola! ¿En qué puedo ayudarte?      │
│                                                 │
│                                                 │
├─────────────────────────────────────────────────┤
│ lfm2.5:8b │ Escribe tu mensaje...              │
├─────────────────────────────────────────────────┤
│ Ctrl+C Salir                                    │
└─────────────────────────────────────────────────┘
```

El selector inferior permite cambiar entre los modelos disponibles en Ollama.

---

# Arquitectura

```text
                    ┌───────────────┐
                    │    Textual    │
                    │      TUI      │
                    └───────┬───────┘
                            │
                            ▼
                    ┌───────────────┐
                    │ SimpleMemory  │
                    └───────┬───────┘
                            │
                 ┌──────────┴──────────┐
                 │                     │
                 ▼                     ▼
          memory.json             Ollama Chat
                 │                     │
                 │                     ▼
                 │               Streaming
                 │                     │
                 └─────────────────────┘
                            │
                            ▼
                       RichLog
```

## Flujo de una conversación

```text
Usuario escribe mensaje
          │
          ▼
SimpleMemory.add()
          │
          ├── Guarda mensaje
          │
          ├── ¿Superó el límite?
          │       │
          │      Sí
          │       ▼
          │   Genera resumen
          │
          └── Guarda memory.json
          │
          ▼
SimpleMemory.messages()
          │
          ▼
       Ollama
          │
          ▼
      Respuesta
          │
          ▼
SimpleMemory.add()
          │
          ▼
       RichLog
```

---

# Tecnologías

- **Python**
- **Textual** — interfaz TUI.
- **Ollama** — ejecución de modelos locales.
- **JSON** — almacenamiento persistente de memoria.

---

# Estado del proyecto

Proyecto experimental orientado a crear una interfaz TUI ligera para modelos locales mediante Ollama.

La arquitectura está pensada para poder extenderse posteriormente con:

- Herramientas para el modelo.
- Memoria más avanzada.
- Diferentes tipos de interacción.
- Gestión de contexto.
- Agentes capaces de utilizar herramientas externas.
