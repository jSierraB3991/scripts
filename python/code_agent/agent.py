import ollama
import json
import libs
from tools_descript import my_tools, available_tools

class Agent:
    def __init__(self):
        self.messages = [{
            "role": "system",
            "content": """
                Eres un agente de código.
                Tu trabajo es ayudar al usuario a modificar y crear proyectos.

                Tu tienes acceso a herramientas para:
                - listar directorios
                - leer archivos
                - escribir archivos

                Antesde modificar un archivo, léelo primero cuando sea necesario para comprender su contenido.
                No inventes  el contenido de archivos que no hayas leído.
                Trabaja únicamente con las herramientas disponibles
            """
        }]

    def run_tools(self, assistant_tools_calls):
        for tool_call in assistant_tools_calls:
            function_name = tool_call["function"]["name"]
            arguments = tool_call["function"]["arguments"]
            print(
                f"\n[TOOL {function_name}]"
                f"{json.dumps(arguments, ensure_ascii=False)}"
            )
            function = available_tools.get(function_name)
            if function is None:
                result = f"Tool desconocida {function_name}"
            else:
                try:
                    result = function(**arguments)
                except Exception as e:
                    result = f"Error ejecutando tool {e}"
            print("[RESULTADO]")
            print(result)
            self.messages.append({
                "role": "tool",
                "content": result
            })

    def run_agent(self, user_input:str):
        self.messages.append({
            "role": "user",
            "content": user_input,
        })

        while True:
            response = ollama.chat(model=libs.MODEL, messages=self.messages, tools=my_tools)
            asistant_messages = response.message
            self.messages.append(asistant_messages)

            if not asistant_messages.get("tool_calls"):
                print("\nAgent:")
                print(asistant_messages.get("content", ""))
                break

            self.run_tools(asistant_messages.tool_calls)

    def prepare(self):
        while True:
            try:
                user_input = input("> ")
            except KeyboardInterrupt:
                print()
                break

            if user_input.lower() in ("exit", "quit"):
                break
            if not user_input.strip():
                continue
            self.run_agent(user_input.strip())
