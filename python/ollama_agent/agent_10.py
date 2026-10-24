import uuid
from typing import TypedDict, Annotated, Literal

from pydantic import BaseModel, Field

from langchain.chat_models import init_chat_model
from langgraph.graph import StateGraph, START, END
from langgraph.checkpoint.memory  import InMemorySaver
from langgraph.graph.message import add_messages

#MODEL = "ollama:lfm2.5-thinking:1.2b"
MODEL = "openrouter:nvidia/nemotron-3.5-super-vl:free"

llm = init_chat_model(MODEL)

class IntentClassifier(BaseModel):
    message_intent: Literal['chat', 'knowledge', 'code'] = Field(..., description='Classify the yser wants to just chat, ask for knowloedge or change in the project')

class State(TypedDict):
    messages: Annotated[list, add_messages]
    message_intent: str | None

def classify_intent(state: State):
    structured_llm = llm.with_structured_output(IntentClassifier)
    result = structured_llm.invoke([
        {'role': 'system', 'content': 'Determine/classify whether the user wants to chat ("chat"), retrive knowlodge ("knowledge") or change code ("code")'},
        { 'role': 'user', 'content': state['messages'][-1].content}
    ])
    return { 'message_intent': result.message_intent}

def prompt_llm_chat(state: State):
    messages= [{'role': 'system', 'content': 'You are a talkative chatbot for a fun. Be nice'}] + state['messages']
    response = llm.invoke(messages)
    return  { 'messages': [{'role': 'assistant', 'content': response.content }]}

def prompt_llm_rag(state: State):
    messages= [{'role': 'system', 'content': 'No matter what the user say, always say "I am the RAG agent"'}] + state['messages']
    response = llm.invoke(messages)
    return  { 'messages': [{'role': 'assistant', 'content': response.content }]}

def prompt_llm_code(state: State):
    messages= [{'role': 'system', 'content': 'No matter what the user say, always say "I am the CODING agent"'}] + state['messages']
    response = llm.invoke(messages)
    return  { 'messages': [{'role': 'assistant', 'content': response.content }]}


def main():
    grhap_builder = StateGraph(State)
    grhap_builder.add_node('classifier', classify_intent)
    grhap_builder.add_node('chat_agent', prompt_llm_chat)
    grhap_builder.add_node('rag_agent', prompt_llm_rag)
    grhap_builder.add_node('coding_agent', prompt_llm_code)

    grhap_builder.add_edge(START, 'classifier')
    grhap_builder.add_conditional_edges('classifier', lambda state: state['message_intent'], {'chat': 'chat_agent', 'knowledge': 'rag_agent', 'code': 'coding_agent'})

    grhap_builder.add_edge('chat_agent', END)
    grhap_builder.add_edge('rag_agent', END)
    grhap_builder.add_edge('coding_agent', END)

    checkpointer = InMemorySaver()

    ghrap = grhap_builder.compile(checkpointer=checkpointer)
    config = { 'configurable': { 'thread_id': uuid.uuid4() } }
    while True:
        user_message = input(f"{MODEL} > ")
        result = ghrap.invoke({'messages': [{'role': 'user', 'content': user_message}]}, config=config)
        print(result['messages'][-1].content)

if __name__ == "__main__":
    main()