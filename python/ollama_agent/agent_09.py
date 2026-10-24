import uuid

from langchain.chat_models import init_chat_model
from langgraph.graph import MessagesState, StateGraph, START, END
from langgraph.checkpoint.memory  import InMemorySaver

MODEL = "ollama:qwen2.5-coder:3b"
#MODEL = "openrouter:nvidia/nemotron-3.5-super-vl:free"

llm = init_chat_model(MODEL)

def prompt_llm(state: MessagesState):
    response = llm.invoke(state["messages"])
    return {'messages': [response]}

def main():
    ghrap_builder = StateGraph(MessagesState)
    ghrap_builder.add_node(prompt_llm)
    ghrap_builder.add_edge(START, 'prompt_llm')
    ghrap_builder.add_edge('prompt_llm', END)

    checkpointer = InMemorySaver()

    ghrap = ghrap_builder.compile(checkpointer=checkpointer)
    config = { 'configurable': { 'thread_id': uuid.uuid4() } }
    while True:
        user_message = input("Enter message: ")
        print(ghrap.invoke({'messages': [{'role': 'user', 'content': user_message}]}, config=config))

if __name__ == "__main__":
    main()