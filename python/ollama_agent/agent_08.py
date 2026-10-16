
from langchain_openrouter import ChatOpenRouter
from langchain_core.messages import SystemMessage, HumanMessage
from dotenv import load_dotenv

load_dotenv()
model = ChatOpenRouter(
    model="nvidia/nemotron-3.5-super-vl:free",
    temperature=0,
    max_retries=1,
)

messages = [
    SystemMessage(
        content="Eres un asistente experto en Go. Responde de forma concisa."
    ),
    HumanMessage(
        content="¿Cuándo debería usar un channel en Go?"
    ),
]

response = model.invoke(messages)

print(response.content)
