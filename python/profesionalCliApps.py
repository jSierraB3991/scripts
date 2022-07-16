#! /usr/bin/python3

import click

@click.command()
@click.option("--name", prompt = "Enter you name", help="The name of the user")
def hello(name):
    click.echo(f"Hello {name}!")


if __name__ == "__main__":
    hello()
