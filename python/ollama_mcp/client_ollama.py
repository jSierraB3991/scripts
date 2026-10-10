
import asyncio
import json

from mcp import ClientSession, StdioServerParameters
from mcp.client.stdio import stdio_client
from ollama import chat
from pathlib import Path



MODEL = "qwen3.5:9b"

SERVER = StdioServerParameters(
    command=".venv/bin/python",
    args=["server.py"],
)


def result_to_text(result) -> str:
    parts = [
        item.text
        for item in result.content
        if hasattr(item, "text")
    ]
    return "\n".join(parts)


async def main():
    question = input("¿Qué quieres consultar? ")

    async with stdio_client(SERVER) as (read, write):
        async with ClientSession(read, write) as session:
            await session.initialize()

            # 1. Descubrir las herramientas MCP.
            discovered = await session.list_tools()

            print("\nHerramientas MCP disponibles:")
            for tool in discovered.tools:
                print(f"- {tool.name}: {tool.description}")

            # 2. Convertir el esquema MCP al formato de Ollama.
            ollama_tools = [
                {
                    "type": "function",
                    "function": {
                        "name": tool.name,
                        "description": tool.description or "",
                        "parameters": tool.inputSchema,
                    },
                }
                for tool in discovered.tools
            ]

            messages = [
                {
                    "role": "system",
                    "content": f"The home of user is {Path.home()}",
                },
                {
                    "role": "user",
                    "content": question,
                }
            ]

            # 3. Ciclo de llamadas a herramientas.
            for _ in range(5):
                response = chat(
                    model=MODEL,
                    messages=messages,
                    tools=ollama_tools,
                )

                message = response.message
                messages.append(message)

                if not message.tool_calls:
                    print("\nRespuesta:")
                    print(message.content)
                    break

                # 4. Ejecutar las herramientas que pide el modelo.
                for call in message.tool_calls:
                    name = call.function.name
                    arguments = call.function.arguments

                    print(f"\nMCP -> {name}({arguments})")

                    result = await session.call_tool(
                        name,
                        arguments=arguments,
                    )

                    output = result_to_text(result)

                    messages.append({
                        "role": "tool",
                        "tool_name": name,
                        "content": output,
                    })
            else:
                print("Se alcanzó el límite de iteraciones.")


if __name__ == "__main__":
    asyncio.run(main())
