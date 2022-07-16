#! /usr/bin/python3

import click


@click.group
def my_commands():
    pass


@click.command()
@click.option("--name", prompt = "Enter you name", help="The name of the user")
def hello(name):
    click.echo(f"Hello {name}!")

PRIORITIES = {
        "o": "OPTIONAL",
        "l": "Lower",
        "m": "MEDIUM",
        "h": "High",
        "c": "CRITICAL"
}

@click.command()
@click.argument("priority", type=click.Choice(PRIORITIES.keys()), default="m")
@click.argument("file", type=click.Path(exists=False), required=0)
@click.option("-n", "--name", prompt="Enter the TODO name", help="The name of the todo item")
@click.option("-d", "--desc", prompt="Describe the TODO", help="The description of the todo item")
def add_todo(name, desc, priority, file):
    file_name = file if file is not None else "myTodo.txt"
    with open(file_name, "a+") as f:
        f.write(f"{name}: {desc} [Priority: {PRIORITIES[priority]}]\n")
        


@click.command()
@click.argument("index", type=int, required=1)
def delete_todo(index):
    with open("myTodo.txt", "r") as f:
        todo_list = f.read().splitlines()
        todo_list.pop(index)

    with open("myTodo.txt", "w") as f:
        f.write("\n".join(todo_list))
        f.write('\n')


@click.command
@click.option("-p","--priority", type=click.Choice(PRIORITIES.keys()))
@click.argument("file", type=click.Path(exists=True), required=0)
def list_todos(priority, file):
    file_name = file if file is not None else "myTodo.txt"
    with open(file_name, "r") as f:
        todo_list = f.read().splitlines()
    if priority is None:
        for id, todo in enumerate(todo_list):
            print(f"({id}) - {todo}")
    else:
        for ixd, todo in enumerate(todo_list):
            if f"[priority: {PRIORITIES[priority]}]" in todo:
                print(f"({idx}) - {todo}")


my_commands.add_command(hello)
my_commands.add_command(add_todo)
my_commands.add_command(delete_todo)
my_commands.add_command(list_todos)

if __name__ == "__main__":
    my_commands()
