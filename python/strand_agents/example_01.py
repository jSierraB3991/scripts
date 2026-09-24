from strands import Agent
from strands.models.ollama import OllamaModel
from strands.vended_tools import file_editor
from strands.memory import MemoryManager
from strands.vended_memory_stores.test_memory_store import TestMemoryStore
from strands import tool

@tool
def letter_counter(word: str, letter: str) -> int:
    """
    Count occurrences of a specific letter in a word.

    Args:
        word (str): The input word to search in
        letter (str): The specific letter to count

    Returns:
        int: The number of occurrences of the letter in the word
    """
    if len(letter) != 1:
        raise ValueError("The 'letter' parameter must be a single character")

    return word.lower().count(letter.lower())

# Persists to ~/.strands/memory/notes.json by default. Survives restarts.
store = TestMemoryStore(name="notes")

# Ephemeral: nothing is written to disk, and a fresh instance forgets everything.
#scratch = TestMemoryStore(name="notes", persist=False)
# Explicit file location instead of the default under ~/.strands/memory/.
#project = TestMemoryStore(name="notes", path="./notes.json")


#from strands.models.gemini import GeminiModel
# Reads GEMINI_API_KEY from the environment.
#model = GeminiModel(model_id="gemini-3.8-flash")

model = OllamaModel(host="http://localhost:11434", model_id="lfm2.5:8b")
agent = Agent(model=model, tools=[letter_counter, file_editor],memory_manager=MemoryManager(stores=[store]))
agent_result = agent('How many letter R\'s are in the word "strawberry"? Write the answer to answer.txt.')
print(agent_result)