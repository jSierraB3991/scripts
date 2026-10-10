
import asyncio
import json
import os

from dotenv import load_dotenv
from openai import OpenAI

from mcp import ClientSession, StdioServerParameters
from mcp.client.stdio import stdio_client


load_dotenv()

MODEL = "gpt-4.1-mini"
client = OpenAI(api_key=os.environ["OPENAI_API_KEY"])

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

            # 1. Descubrir herramientas MCP.
            discovered = await session.list_tools()

            print("\nHerramientas MCP disponibles:")
            for tool in discovered.tools:
                print(f"- {tool.name}: {tool.description}")

            # 2. Convertir el esquema MCP a function calling.
            openai_tools = [
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
                    "role": "user",
                    "content": question,
                }
            ]

            # 3. Ejecutar el ciclo de herramientas.
            for _ in range(5):
                response = client.chat.completions.create(
                    model=MODEL,
                    messages=messages,
                    tools=openai_tools,
                )

                message = response.choices[0].message
                messages.append(message)

                if not message.tool_calls:
                    print("\nRespuesta:")
                    print(message.content)
                    break

                # 4. Ejecutar cada llamada mediante MCP.
                for call in message.tool_calls:
                    name = call.function.name
                    arguments = json.loads(
                        call.function.arguments
                    )

                    print(f"\nMCP -> {name}({arguments})")

                    result = await session.call_tool(
                        name,
                        arguments=arguments,
                    )

                    output = result_to_text(result)

                    messages.append({
                        "role": "tool",
                        "tool_call_id": call.id,
                        "content": output,
                    })
            else:
                print("Se alcanzó el límite de iteraciones.")


if __name__ == "__main__":
    asyncio.run(main())
