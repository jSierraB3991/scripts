import streamlit as st
from langchain_ollama import ChatOllama
from langchain_core.tools import tool
from langchain_core.messages import  HumanMessage, AIMessage
from langchain.agents import create_agent
from langgraph.checkpoint.sqlite import SqliteSaver

@tool(name_or_callable="calculadora", description="Úsala para resolver operaciones matemáticas simples. Ej: '5 * (2+3)'")
def calculadora(query: str) -> str:
    try:
        result = eval(query)
        return f"El resultado de {query} es {result}"
    except Exception:
        return "Lo siento, no se pudo resolver está operación"

def mostrar_historial(thread_id: str, db_path: str = "checkpoints.sqlite"):
    """Lee el historial de una conversación directo de sqlite, sin tocar el agente."""
    with SqliteSaver.from_conn_string(db_path) as checkpointer:
        config = {"configurable": {"thread_id": thread_id}}
        checkpoint = checkpointer.get(config)
        if not checkpoint:
            return []
        return checkpoint["channel_values"].get("messages", [])

THREAD_ID="conversacion-1"
def main():
    chat = ChatOllama(model="gemma4:12b")

    prompt = """Eres *orion*, un agente conversacional en español, con memoria y capacidad de razonar.\n
                        Respondes de forma clara, cercana y didáctica.\n
                        Puedes usar herramientas como una calculadora para ayudar a mejorar al usuario.\n
                        Recuerda lo que el usuario te ha dicho antes y sé útil, empático y profesional"""
    st.set_page_config(page_title="Orian - Agent IA LOCAL", layout="centered")
    st.title("Orion agente de IA")
    st.markdown("Hola, Soy orion")
    if "history" not in st.session_state:
        st.session_state.history = mostrar_historial(thread_id=THREAD_ID)
    for msg in st.session_state.history:
        with st.chat_message("user" if isinstance(msg, HumanMessage) else "assistant"):
            st.markdown(msg.content)
    with SqliteSaver.from_conn_string("checkpoints.sqlite") as checkpointer:
        config = {"configurable": {"thread_id": THREAD_ID}}
        user_input = st.chat_input("Escribe tu mensaje: ")
        if user_input:
            try:
                st.session_state.history.append(HumanMessage(content=user_input))
                agent = create_agent(chat, tools=[calculadora], system_prompt=prompt, checkpointer=checkpointer)
                result = agent.invoke(
                    {"messages": [("user", user_input)]}, config=config
                )
                resp_final = result["messages"][-1].content
                st.session_state.history.append(AIMessage(content=resp_final))
                with st.chat_message("assistant"):
                    st.markdown(resp_final)
            except Exception as e:
                print(e)


        
if __name__ == "__main__":
    main()
