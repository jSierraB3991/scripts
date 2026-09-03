from ollama import chat, Client,ChatResponse
import libs 
import agent

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
    agent.Agent().prepare()

if __name__ == "__main__":
    main()