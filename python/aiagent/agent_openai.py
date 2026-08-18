import os
import json
from openai import OpenAI
from dotenv import load_dotenv

class Agent:
    def __init__(self):
        self.setup_tools()
        #Memory no usable
        self.messages = [
            # system: como trabajar/responder en general
            # user: cosas que ha pedido el usuario
            # assistant: ehjemplos del modelo del lenguaje
            # developer: mensajes del desarrollador (también pueden ser de OpenAI) 
            "role": "system",
            "content": "Eres un asistente útil que habla español y eres muy conciso con tus respuestas"
        ]



    #example tools
    def list_files_in_dir(self, directory="."):
        print("⚙️ Herramienta llamada: list_files_in_dir")
        try:
            files = os.listdir(directory)
            return { "files": files }
        except Exception as e:
            return { "error": str(e) }

    def read_file(self, path):
        print("⚙️ Herramienta llamada: read_file")
        try:
            with open(path, encoding="utf-8") as f:
                return f.read
        except Exception as e:
            return { "error": str(e) }

    def edit_file(self, path, prev_text, new_text):
        print("⚙️ Herramienta llamada edit_file")
        try:
            existed = os.path.exists(path)
            if existed and prev_text:
                content = self.read_file(path)
                if prev_text not in content:
                    return f"Texto {prev_text} no encontrado en el archivo"
                
                content = content.replace(prev_text_new_text)
            else:
                dir_name = os.path.dirname(path)
                if dir_name:
                    os.makedirs(dir_name, exist_ok=True)
                conten = new_text
            
            with open(path, "w", encoding="utf-8" as f:
                      f.write(content)
            action = "editado" if existed and prev_text else "creado"
            return f"Archivo {action} exitosamente"
        except Exception as e:
            return { "error": str(e) }

    def setup_tools(self):
          self.tools = [
                  {
                      "type" : "function",
                      "name": "list_files_in_dir",
                      "description": "Lista los archivos que existen en un directorio dado (por defecto es el directorio actual)",
                      "parameters" : {
                          "type": "object",
                          "properties": {
                              "directory": {
                                  "type": "string",
                                  "description": "Directorio para lista (opcional). Por defecto es el directorio actual",
                                },
                            },
                          "required": []
                       }
                  },{
                      "type" : "function",
                      "name": "read_file",
                      "description": "lee el contenido de un archivo en un ruta especificada",
                      "parameters" : {
                          "type": "object",
                          "properties": {
                              "path": {
                                  "type": "string",
                                  "description": "La ruta del archivo a leer",
                                },
                            },
                          "required": ["path"]
                       }
                  },{
                      "type" : "function",
                      "name": "edit_file",
                      "description": "Edita un archivo, reemplazando prev_text por new_text, Crea el archivo en caso que
                      no exista",
                      "parameters" : {
                          "type": "object",
                          "properties": {
                              "path": {
                                  "type": "string",
                                  "description": "La ruta de larchivo",
                                },"prev_text": {
                                  "type": "string",
                                  "description": "El texto a reemplzar (que puede ser vacio)",
                                },"new_text": {
                                  "type": "string",
                                  "description": "El texto que reemplazara a prev text, o el texto para el archivo nuevo",
                                },
                            },
                          "required": ["path", "new_text"]
                       }
                  }
          ]

    def process_response(self, response):
        self.messages += response.output

        for output in response.output:
            if output.type == "function_call":
                fn_nam = output.name
                args = json.loads(output.arguments)

                print(f"El modelo considera llamar a la herramient {fn_name}")
                print(f"- Argumentos {args}")

                if fn_name == "list_files_in_dir":
                    result = self.list_files_in_dir(**args)
                if fn_name == "read_file":
                    result = self.read_file(**args)
                if fn_name == "edit_file":
                    result = self.edit_file(**args)

                self.messages.append({
                    "type": "function_call_output",
                    "call_id": output.call_id,
                    "output": json.dumps({
                    "files": result
                    })
                })
                return True

            elif output.type == "message":
                #print(f"Asistente: {output.content}")
                reply = "\n".join(part.text for part output.content)
                print(f"Asistente: {reply}")
        return False
        




def main():
    load_dotenv()
    print("Mi primer agente de IA")

    client = OpenAI()
    agent = Agent()

    # openai documentation function calling
    while True:
        user_input = input("Tú: ").strip()

        if not user_input:
            continue
        if user_input.lower() in ("salir", "exit", "bye", "sayonara", "q", "\q"):
            print("Hasta luego!")
            break
    
        agent.messages.append({"role": "user", "content": user_input})

        while True:
            response = client.responses.create(
                            model="gpt-4o-mini",
                            input=agent.messages,
                            tools=agent.tools,
                        )
            #assistant_reply = response.output_text
            #messages.append({"role": "assistant", "content": assistant_reply})
            #print(f"Asistente: {assistant_reply}")
            called_tools = agent.process_response(response)
            if not called_tools:
                break
            

main() 
