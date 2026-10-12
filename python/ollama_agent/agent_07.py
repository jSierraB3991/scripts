from base64 import b64encode
from dataclasses import dataclass
import requests
import os

from dotenv import load_dotenv

from langchain.agents import create_agent
from langchain.agents.structured_output import ToolStrategy
from langchain.chat_models import init_chat_model
from langchain.messages import AIMessage, HumanMessage, SystemMessage
from langchain.tools import tool, ToolRuntime
from langchain_ollama import ChatOllama

from langgraph.checkpoint.memory import InMemorySaver
from pydantic import BaseModel, Field

API_KEY = os.getenv("OPEN_WATER_KEY")
MODEL_USE = "lfm2.5:8b"
load_dotenv()

@dataclass
class Context:
    user_id: str

@dataclass
class ResponseFormat(BaseModel):
    summary: str = Field(description="Resume el estado del cielo, la temperatura y la sensación térmica.")
    temperature_celcius: float = Field(description="Temperatura en Celsius")
    humedity: float = Field(description="Humedad relativa en porcentaje")

@tool("locate_user", description="Look up a user's city based on context")
def locate_user(runtime: ToolRuntime[Context]):
    print(f"🤖 call locate_user with user id {runtime.context.user_id}")
    match runtime.context.user_id:
        case "ABC123":
            return "Lima"
        case "MY":
            return "Turbaco, Bolivar"
        case "XYZ":
            return "Bogotá D.c"
        case "HJK111":
            return "Buenos Aires"
        case _: 
            return "Unknown"

@tool("get_weather", description="Return weather by latitude and logitude", return_direct=False)
def get_weather(latitude, longitude: float):
    print(f"🤖 call get_weather with: lat {latitude} and lon {longitude}")
    response = requests.get(
        "https://api.openweathermap.org/data/2.5/weather",
        params={
            "lat": latitude,
            "lon": longitude,
            "units": "metric",
            "lang": "es",
            "appid": API_KEY,
        },
        timeout=(50, 60),  # conexión: 5 s, respuesta: 30 s
    )
    response.raise_for_status()
    result = response.json()
    return result

@tool('get_location', description="Return latitude and longitude information for a given city", return_direct=False)
def get_location(city: str):
    print(f"🤖 call get_location with args city: {city}")
    response = requests.get(
        "https://api.openweathermap.org/geo/1.0/direct",
        params={
            "q":  city,
            "appid": API_KEY,
        },
        timeout=(50, 60),  # conexión: 5 s, respuesta: 30 s
    )
    response.raise_for_status()
    return response.json()

def get_weather_by_city():
    modelOllama = ChatOllama(model=MODEL_USE)
    agent = create_agent(
        model=modelOllama,
        tools=[get_weather, get_location],
        system_prompt="You are a helpful weather assitant, who always craks jokes is humorous while remaining helpful",
        temperature=0.1,
    )

    user_input = "What is weather like in Vienna, austria?"# answerme in spanish"
    print(f"💬 user > {user_input}")
    response = agent.invoke({
        "messages": [
            {
                "role": "user",
                "content": user_input
            }
        ]
    })

    print(response)
    print(f"🤖 {MODEL_USE}: {response['messages'][-1].content}")

def chat_model_test():
    #kwargs
    #temperature: Model temperature for controlling randomness.
    #max_tokens: Maximum number of output tokens.
    #timeout: Maximum time (in seconds) to wait for a response.
    #max_retries: Maximum number of retry attempts for failed requests.
    #base_url: Custom API endpoint URL.
    chat_model = init_chat_model(
        model=MODEL_USE,
        model_provider="ollama",
        temperature=0.1,
    )

    conversation = [
        SystemMessage("You are a helpful assistant for question regarding programming"),
        HumanMessage("What is Python?"),
        AIMessage("Python is an interpreted programming language."),
        HumanMessage("When was it released?"),
    ]

    for chunk in chat_model.stream(conversation):
        print(chunk.text, end="", flush=True)
    print("")
    #print(f"id: {response.id}")
    #print(f"response_metadata: {response.response_metadata}")
    #print(f"usage_metadata: {response.usage_metadata}")
    #print(f"🤖 {MODEL_USE}: {response.content}")

def describe_image():
    MODEL_IMAGE = "ministral-3:8b"
    model = init_chat_model(
        model=MODEL_IMAGE,
        model_provider="ollama",
    )
    user_input = "Describe el contenido de esta imagen."
    print(f"user> {user_input}")
    with open("image.jpg", "rb") as image:
        image_b64 = b64encode(image.read()).decode("utf-8")
    messages = HumanMessage(
            content = [
                {
                    "type": "text",
                    "text": user_input
                },
                {
                    "type": "image",
                    "base64": image_b64,
                    "mime_type": "image/jpeg",
                },
            ])

    response = model.invoke([messages])
    print(f"🤖 {MODEL_USE}: {response.content}")

chat_model = init_chat_model(
    model=MODEL_USE,
    model_provider="ollama",
    temperature=0.3,
)

checkpointer = InMemorySaver()

agent = create_agent(
    model=chat_model,
    tools=[get_weather, get_location, locate_user],
    system_prompt="You are a helpful weather assistant who uses locate_user to determine their city; do not make up the weather — use the available tools. You always crack jokes and maintain a good sense of humor while remaining helpful.",
    context_schema=Context,
    response_format=ToolStrategy(ResponseFormat),
    checkpointer=checkpointer,
)

config = { 'configurable': { 'thread_id': "conversation01" }}
user_input = "What is the weather like?, not invent the city?"# answer me in spanish"
print(f"💬 user > {user_input}")
response = agent.invoke({
    "messages": [
        {
            "role": "user",
            "content": user_input
        }
    ]}, 
    config = config,
    context = Context(user_id="MY")
)

#print(response)
structured = response.get("structured_response")
if structured:
    print(f"""
        🤖 {MODEL_USE}:
        🌤️ {structured.summary}

        🌡️ Temperatura: {structured.temperature_celcius:.1f} °C
        💧 Humedad:     {structured.humedity:.0f}%
        """.strip())
else:
    messages = response.get("messages")
    print(f"🤖 {MODEL_USE}: {messages[-1].content}")
