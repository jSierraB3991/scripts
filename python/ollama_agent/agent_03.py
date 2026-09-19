import os
from pathlib import Path
from langchain_ollama import ChatOllama
from langchain.agents import create_agent
from langchain_core.tools import tool
from langchain_core.prompts import PromptTemplate

# --- PASO 1: Definir la Herramienta (Tool) ---
# Es vital que el docstring sea claro, ya que el modelo lee esto 
# para decidir CUÁNDO usar la herramienta.
@tool
def create_file_tool(filename: str, content: str) -> str:
    """
    Útil para crear un archivo de texto en el sistema. 
    El usuario debe proporcionar el nombre del archivo y el contenido.
    """
    try:
        with open(filename, "w", encoding="utf-8") as f:
            f.write(content)
        return f"Éxito: El archivo '{filename}' ha sido creado correctamente."
    except Exception as e:
        return f"Error al crear el archivo: {str(e)}"

# --- PASO 2: Configurar el Modelo ---
# Usamos llama3 (o el modelo que tengas en Ollama)
llm = ChatOllama(model="gemma4:12b")

# --- PASO 3: Configurar el Prompt del Agente ---
# El prompt debe indicar al modelo que tiene herramientas disponibles.
prompt = """Eres un asistente útil que puede usar herramientas para realizar tareas.
Tienes las siguientes herramientas a tu disposición:

[create_file_tool]

Usa el siguiente formato para responder:
Thought1: Do I need to use a tool? Yes
Action: {action}
Action Input: {action_input}
Observation: (La respuesta de la herramienta)
... (esto puede repetirse si es necesario)

Thought1: I have enough information to answer.
Final Answer: (Tu respuesta final al usuario)

Pregunta del usuario: {input}
{agent_scratchpad}"""


# --- PASO 4: Crear el Agente ---
# En este ejemplo, simplificamos la lógica para que sea directo.
tools = [create_file_tool]

# --- PASO 5: Ejecutar el Agente ---
executor = create_agent(
    model=llm, 
    tools=tools, 
    system_prompt=prompt,
)

# Prueba del sistema
if __name__ == "__main__":
    query = "dentro de la carpeta 'basic_agent' en este mismo path, crea tres archivos, un archivo para el main para iniciar el proyecto, pedirle el nombre al usuario y pedirle el nombre del agente de código, uno llamado agente que tenga todo lo que un agente que tiene el bucle del agente y uno de tools, donde estén las tools que puede usar el agente.\nEl agente se va a hacer con ollama y lang chain, con el modelo gemma4:12b"
    result = executor.invoke(
        {"messages": [("user", query)]}
    )
    resp_final = result["messages"][-1].content
    print(resp_final)