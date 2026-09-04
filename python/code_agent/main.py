from ollama import chat, ChatResponse
import libs 
import agent
from tools.database.get_memory_by_key_tool import get_memory_by_key
from tools.database.save_memory_tool import save_memory

def main():
    print(f"Code Agent - {libs.MODEL}")

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