from progress_bar import  BarraCarga
import ollama
import libs
import json

class SimpleMemory:
    def __init__(self, max_msg = 10, num_summarize=4):
        self.memory = []
        self.max_msg = max_msg
        self.num_summarize = num_summarize
        self.sumarize = ""

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

        chargeBar = BarraCarga()
        chargeBar.iniciar()
        response = ollama.chat(model=libs.MODEL, messages=[{"role": "user", "content": prompt}])
        chargeBar.finalizar()
        return response.message.content or ""

    def save_json(self, file_name: str = "memory.json"):
        data = {
            "summary": self.sumarize,
            "memory": self.memory,
        }
        with open(file_name, "w", encoding="utf-8") as f:
            json.dump(data, f, ensure_ascii=False, indent=4)
        pass