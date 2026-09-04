import ollama
import json
import libs
from tools_descript import my_tools, available_tools,tools_with_question
from barra_carga import  BarraCarga

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
                Trabaja únicamente con las herramientas disponibles.
                Solo búsca archivos dentro de la carpeta donde estás, nunca busques fuera.
            """
        }]

    def validate_tool(self, function_name: str, arguments: any) -> bool:
        if function_name not in tools_with_question:
            print(
                f"\n[TOOL {function_name}]"
                f"{json.dumps(arguments, ensure_ascii=False)}"
            )   
            return True
        while True:
            verification_tool = f"voy a usar la herramienta: {function_name} con los argumentos: {arguments}  \n está de acuerdo? Y/n: "
            response_user = input(verification_tool)
            if response_user.lower() == "y":
                return True
            else:
                while True:
                    why_user = input("Porque no lo ha verificado?: ")
                    if why_user.strip() != "":
                        self.messages.append({
                            "role": "user",
                            "content": f"El usuario no permitio la ejecución de la por: {why_user}"
                        })
                        break
                break
        pass
        return False


    def run_tools(self, assistant_tools_calls):
        for tool_call in assistant_tools_calls:
            function_name = tool_call["function"]["name"]
            arguments = tool_call["function"]["arguments"]

            pass_user = self.validate_tool(function_name=function_name, arguments=arguments)
            if not pass_user:
                continue
                         
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
            chargeBar = BarraCarga()
            chargeBar.iniciar()
            response = ollama.chat(model=libs.MODEL, messages=self.messages, tools=my_tools)
            chargeBar.finalizar()
            asistant_messages = response.message
            self.messages.append(asistant_messages)
            if asistant_messages.thinking.strip() != "":
                print(f"pensando: {asistant_messages.thinking.strip()}")

            if not asistant_messages.get("tool_calls"):
                print("\nAgent:")
                print(asistant_messages.get("content", ""))
                break

            self.run_tools(asistant_messages.tool_calls)

    def print_bye(self, message="adiós", tipo_color=libs.RED):
        emoji = "👋"
        print(f"{tipo_color}{libs.YELLOW} {emoji} {message} {libs.RESET}")
        pass

    def prepare(self, name_user: str):
        if name_user != "":
            print(f"Hola {name_user} como puedo ayudarlo hoy?")
        while True:
            try:
                user_input = input(f"{name_user} > ")
            except KeyboardInterrupt:
                self.print_bye("Saliendo por interrupción")
                break

            if user_input.lower() in ("exit", "quit"):
                self.print_bye()
                break
            if not user_input.strip():
                continue
            self.run_agent(user_input.strip())
