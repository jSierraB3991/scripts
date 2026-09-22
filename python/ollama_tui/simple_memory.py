import json
import ollama
DEFAULT_MODEL = "lfm2.5:8b"

class SimpleMemory:
    def __init__(self, max_msg = 12, num_summarize=8):
        self.memory = []
        self.max_msg = max_msg
        self.num_summarize = num_summarize
        self.sumarize = ""
        self.load_data_file()

    def load_data_file(self, ruta_archivo: str = "memory.json"):
        """Lee el archivo JSON especificado y asigna los datos a self.sumarize y self.memory."""
        try:
            with open(ruta_archivo, "r", encoding="utf-8") as archivo:
                datos = json.load(archivo)
                # Extraemos los valores usando .get() para evitar errores si no existen las claves
                self.sumarize = datos.get("summary", "")
                self.memory = datos.get("memory", [])

        except FileNotFoundError:
            print(f"Error: El archivo '{ruta_archivo}' no existe.")
        except json.JSONDecodeError:
            print(f"Error: El archivo '{ruta_archivo}' no tiene un formato JSON válido.")

    def add(self, role:str, msg: str):
        self.memory.append({"role": role, "content": msg})
        if len(self.memory) > self.max_msg:
            print(f"Preparando para resumir {self.num_summarize} messages")
            to_summarize = self.memory[:self.num_summarize]
            self.memory = self.memory[self.num_summarize:]
            self.sumarize = self.sumarize_memory(to_summarize)
        print(f"Mensajes actuales en la memoria: {self.memory}")
        self.save_json()

    def messages(self):
        if self.sumarize:
            mem_sumary = {
                "role": "system",
                "content": f"Memorias de conversación: {self.sumarize}"
            }
            return [mem_sumary] + self.memory
        return self.memory

    def sumarize_memory(self, old_memories: list):
        summary_line = f"Resumen anterior: {self.sumarize}\n" if self.sumarize else  "No hay resumen previo\n"
        prompt = f"""Tu tarea es actualizar la memoria resumen del asistente\n
        {summary_line} \n
        Nuevos mensajes a integrar {old_memories}
        Genera un nuevo resumen corto y consolidado que combine ambos,
        manteniendo datos clave (nombre, acuerdos, fechas, directorio de trabajo, estructura del proyecto)\n
        No guardes respuestas del asistente, solo datops importantes del usuario de manera muy concisa y estructurada\n
        Guarda cosas importantes, no detalles raros o curiosos del usuario\n
        El texto del resumen debe ser corto"""

        response = ollama.chat(model=DEFAULT_MODEL, messages=[{"role": "user", "content": prompt}])
        return response.message.content or ""

    def save_json(self, file_name: str = "memory.json"):
        data = {
            "summary": self.sumarize,
            "memory": self.memory,
        }
        with open(file_name, "w", encoding="utf-8") as f:
            json.dump(data, f, ensure_ascii=False, indent=4)
        pass
