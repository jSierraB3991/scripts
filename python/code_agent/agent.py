import ollama
import json
import libs
from tools.tools_descript import my_tools, available_tools,tools_with_question
from progress_bar import  BarraCarga
from tools.database.get_all_memories import get_all_memories

class Agent:
    def __init__(self):
        self.loop = 0
        self.init_memory()
        
    def init_memory(self):
        self.messages = [{
                "role": "system",
                "content": """
                    Eres un agente de código.
                    Tu trabajo es ayudar al usuario a modificar y crear proyectos.

                    Tu tienes acceso a herramientas para:
                    - listar directorios
                    - leer archivos
                    - escribir archivos

                    Antes de modificar un archivo, léelo primero cuando sea necesario para comprender su contenido.
                    No inventes  el contenido de archivos que no hayas leído.
                    Trabaja únicamente con las herramientas disponibles.
                    A menos que te diga lo contrario, solo búsca archivos dentro de la carpeta raíz del proyecto, nunca fuera.
                    No listes archivos porque sí, si ya lso tienes en memoria, están bien.
                    En python no crees archivo __init__.py, en ninguna parte del código.
                    Tú memoria se vacea varias veces, guarda lo que necesites y creas que puedes necesitar después
                """
            }]
        memories = get_all_memories()
        self.run_agent(f"Estos datos debes tenerlos en memoria, pero no es necesario guardalos, ya están guardados: {'\n'.join(mem.to_model() for mem in memories)}")
    def restart_memory(self, user_input: str):
        print("Limpiando la memoria")
        self.loop = 0
        self.init_memory()
        print(f"Reiniciando consulta del usuario {user_input}")
        self.messages.append({
            "role": libs.ROL_USER,
            "content": user_input,
        })

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
                            "content": f"Error: El usuario no permitio la ejecución de la tool {function_name} por: '{why_user}'"
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
            self.messages.append({
                "role": libs.ROL_TOOL,
                "content": result
            })

    def run_agent(self, user_input:str):
        self.messages.append({
            "role": libs.ROL_USER,
            "content": user_input,
        })

        while True:
            if self.loop > libs.MAX_LOOP_AGENT:
                self.restart_memory(user_input)
            chargeBar = BarraCarga()
            chargeBar.iniciar()
            response = ollama.chat(model=libs.MODEL, messages=self.messages, tools=my_tools)
            chargeBar.finalizar()
            asistant_messages = response.message
            
            self.loop +=1
            if  asistant_messages.thinking:
                self.messages.append({
                    "role": libs.ROL_AGENT,
                    "content": asistant_messages.thinking
                })
            
            if asistant_messages.get("tool_calls"):
                self.run_tools(asistant_messages.tool_calls)
                continue


            if  asistant_messages.content == "":
                print(asistant_messages.role, asistant_messages.thinking if asistant_messages.thinking == None else asistant_messages.thinking.strip())
                continue

            self.loop = 0
            print("\nAgent:")
            print(asistant_messages.content if asistant_messages != "" else asistant_messages)
            break


    def print_bye(self, message="adiós", tipo_color=libs.RED):
        emoji = "👋"
        print(f"{tipo_color}{libs.YELLOW} {emoji} {message} {libs.RESET}")
        pass

    def prepare(self, name_user: str):
        #self.run_agent("Hola, soy un desarollador (tú sabes mi nombre) y quiero que sepas la estructura del proyecto que vamos a contruir, Si no sabe cuál es la estructura del proyecto estructura completa del proyecto con archivos y carpetas en cada nivel, puedes buscarla el la database con la key 'key_structure_proyecto' de no encontrarla, y si tampoco la encuentras, lista todos los archivos y agregala con los archivos y carpetas del proyecto. Despúes de esto, presentate y saludame")
        print("Escribe 'exit' para salir.")
        while True:
            try:
                user_input = input(f"{name_user} - {libs.MODEL} > ")
            except KeyboardInterrupt:
                self.print_bye("Saliendo por interrupción")
                break

            if user_input.lower() in ("exit", "quit"):
                self.print_bye(f"adiós {name_user}")
                break
            if user_input.lower() in ("clear", "cls"):
                libs.clear_screen()
                continue
            if not user_input.strip():
                continue
            self.run_agent(user_input.strip())
