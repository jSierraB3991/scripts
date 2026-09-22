from textual.app import App, ComposeResult
from textual.containers import Horizontal
from textual.widgets import Header, Footer, Input, RichLog, Select
from textual.worker import  get_current_worker

from simple_memory import SimpleMemory, DEFAULT_MODEL

import ollama

class OllamaTUI(App):
    CSS = """
    Screen {
        layout: vertical;
    }
    #chat_log {
        height: 1fr;
        border: solid $accent;
    }
    #bottom_bar {
        height: 3;
    }
    #model_select {
        width: 30;
    }
    #prompt_input {
        width: 1fr;
    }
    """

    BINDINGS = [("ctrl+c", "quit", "Salir")]

    def __init__(self):
        super().__init__()
        self.current_model = DEFAULT_MODEL
        self.history: list[dict] = []
        self.simple_memory = SimpleMemory()

    def compose(self) -> ComposeResult:
        yield Header()
        yield RichLog(id="chat_log", wrap=True, markup=True)
        with Horizontal(id="bottom_bar"):
            yield Select(
                options=[(DEFAULT_MODEL, DEFAULT_MODEL)],
                id="model_select",
                value=DEFAULT_MODEL,
                allow_blank=False,
            )
            yield Input(placeholder="Escribe tu mensaje...", id="prompt_input")
        yield Footer()

    def on_mount(self) -> None:
        self.query_one("#prompt_input", Input).focus()
        self.run_worker(self.cargar_modelos, thread=True)

    def cargar_modelos(self) -> None:
        try:
            resp = ollama.list()
            nombres = [m["model"] for m in resp.get("models", [])]
        except Exception as e:
            self.call_from_thread(self.log_error, f"No se pudo listar modelos: {e}")
            return

        if DEFAULT_MODEL not in nombres:
            nombres.insert(0, DEFAULT_MODEL)

        self.call_from_thread(self.actualizar_select, nombres)

    def actualizar_select(self, nombres: list[str]) -> None:
        select = self.query_one("#model_select", Select)
        select.set_options([(n, n) for n in nombres])
        select.value = self.current_model

    def log_error(self, texto: str) -> None:
        self.query_one("#chat_log", RichLog).write(f"[bold red]Error:[/] {texto}")

    def on_select_changed(self, event: Select.Changed) -> None:
        if event.select.id == "model_select" and event.value:
            self.current_model = str(event.value)
            self.query_one("#chat_log", RichLog).write(
                f"[bold yellow]Modelo cambiado a:[/] {self.current_model}"
            )

    def on_input_submitted(self, event: Input.Submitted) -> None:
        texto = event.value.strip()
        if not texto:
            return

        event.input.value = ""
        log = self.query_one("#chat_log", RichLog)
        log.write(f"[bold cyan]Tú:[/] {texto}")

        self.simple_memory.add("user", texto)
        self.run_worker(self.consultar_ollama, thread=True)

    def consultar_ollama(self) -> None:
        worker = get_current_worker()
        log = self.query_one("#chat_log", RichLog)

        self.call_from_thread(
            log.write, f"[dim]({self.current_model} pensando...)[/dim]"
        )

        try:
            respuesta_completa = ""
            stream = ollama.chat(
                model=self.current_model,
                messages=self.simple_memory.messages(),
                stream=True,
            )
            for chunk in stream:
                if worker.is_cancelled:
                    return
                contenido = chunk["message"]["content"]
                respuesta_completa += contenido

            self.simple_memory.add("assistant", respuesta_completa)
            self.call_from_thread(
                log.write, f"[bold green]{self.current_model}:[/] {respuesta_completa}"
            )
        except Exception as e:
            self.call_from_thread(self.log_error, str(e))


if __name__ == "__main__":
    OllamaTUI().run()