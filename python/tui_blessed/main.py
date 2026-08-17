#!/usr/bin/env python3
"""TUI de tareas (todo) usando blessed."""

from blessed import Terminal

term = Terminal()

tasks = []  # cada tarea: {"text": str, "done": bool}
selected = 0
input_mode = False
input_buffer = ""


def draw():
    print(term.home + term.clear)
    print(term.bold_underline("TODO TUI (blessed)"))
    print(term.dim("↑/↓ mover  |  espacio: completar  |  a: agregar  |  d: eliminar  |  q: salir"))
    print()

    if not tasks:
        print(term.dim("  (sin tareas, presiona 'a' para agregar)"))
    else:
        for i, task in enumerate(tasks):
            box = "[x]" if task["done"] else "[ ]"
            line = f"{box} {task['text']}"
            if task["done"]:
                line = term.strikethrough(line)
            if i == selected:
                print(term.reverse(f"> {line}"))
            else:
                print(f"  {line}")

    print()
    if input_mode:
        print(term.bold("Nueva tarea: ") + input_buffer + "█")


def main():
    global selected, input_mode, input_buffer

    with term.fullscreen(), term.cbreak(), term.hidden_cursor():
        draw()
        while True:
            key = term.inkey()

            if input_mode:
                if key.name == "KEY_ENTER":
                    if input_buffer.strip():
                        tasks.append({"text": input_buffer.strip(), "done": False})
                    input_buffer = ""
                    input_mode = False
                elif key.name == "KEY_ESCAPE":
                    input_buffer = ""
                    input_mode = False
                elif key.name == "KEY_BACKSPACE":
                    input_buffer = input_buffer[:-1]
                elif key and not key.is_sequence:
                    input_buffer += key
            else:
                if key == "q":
                    break
                elif key.name == "KEY_UP" and tasks:
                    selected = (selected - 1) % len(tasks)
                elif key.name == "KEY_DOWN" and tasks:
                    selected = (selected + 1) % len(tasks)
                elif key == " " and tasks:
                    tasks[selected]["done"] = not tasks[selected]["done"]
                elif key == "a":
                    input_mode = True
                elif key == "d" and tasks:
                    tasks.pop(selected)
                    if selected >= len(tasks):
                        selected = max(0, len(tasks) - 1)

            draw()


if __name__ == "__main__":
    main()
