from ollama import chat, Client,ChatResponse
import libs 
import agent
from tools.tool_db import save_memory, get_memory_by_key

def no_use():
    response: ChatResponse = chat(model="qwen2.5-coder:3b", messages=[
        {
            'role': 'user',
            'content': ''
       },
    ])
    print(response['message']['content'])
    print(response.message.content)

def main():
    print(f"Code Agent - {libs.MODEL}")
    print("Escribe 'exit' para salir.\n")

    key_name = get_memory_by_key(libs.KEY_NAME)
    if not key_name:
        name = input("Buenas, como te llamas? ")
        if name.strip() != "":
            save_memory(key=libs.KEY_NAME,content=name.strip(), role_agent=libs.ROL_USER)
    else:
        name = key_name.content

    agent.Agent().prepare(name.strip())

if __name__ == "__main__":
    main()