import sqlite3

from strands import Agent
from strands.models.ollama import OllamaModel
from strands import tool

DB_PATH = "./buy_data.db"
MODEL = "qwen3.5:9b"

@tool
def get_structuct_table_tool(table_name: str) -> any:
    """
    GET table structured

    Args:
        table_name: str -> Name of table 
    Returns:
        any -> Return the table structure or if this table no exists
    """

    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute(f"SELECT name FROM pragma_table_info('{table_name}');")
    rows = cursor.fetchall()

    conn.close()
    names = [r[0] for r in rows]
    return names

@tool
def run_sqlite_query_tool(query: str) -> any:
    """
    RUn Query in the database

    Args:
        query: str: Query of sql t oexecute in the database 
    Returns:
        any: Return the database table or why user stop the query
    """

    if any(word in query.lower() for word in ['update', 'delete', 'drop']):
        print(f"The agent wants to run the following SQL query:\n{query}\n")
        user_answer = input("Are you sure you want to run this query?\n1) Yes\n2) No\nOther) No\n> ")
        if user_answer != "1":
            why_stop = input("Why do you want to stop running the query?\n> ")
            return why_stop

    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute(query)
    rows = cursor.fetchall()

    conn.close()
    return rows

@tool
def get_tables_tool() -> list[str]:
    """
    Get all tables on the database

    Returns:
        list[str]: Return the names of the tables in the database
    """

    conn = sqlite3.connect(DB_PATH)
    cursor = conn.cursor()
    
    cursor.execute("SELECT name FROM sqlite_master WHERE type = 'table';")
    rows = cursor.fetchall()

    conn.close()
    names = [{"table": r[0], "columns": ",".join(get_structuct_table_tool(r[0]))} for r in rows]
    return names

model = OllamaModel(host="http://localhost:11434", model_id=MODEL)
agent = Agent(model=model, 
            tools=[get_tables_tool, run_sqlite_query_tool],
            system_prompt = """You are a database query agent.
                Always retrieve information from the database when answering questions related to the database.
                Before providing the final answer, review the retrieved data for inconsistencies, missing information, or unusual values.
                If you find something unusual or inconsistent, clearly inform the user about it.
                STRICT RULES:
                1. You ONLY answer questions that can be answered using information from the database.
                2. For EVERY database-related question, you MUST use the database tools.
                3. NEVER answer general knowledge, programming, technical, or conversational questions.
                4. If the user's question is NOT related to the database, DO NOT explain, teach, or provide any other information.
                5. For non-database questions, respond EXACTLY with:
                'No data available for this information.'
                6. If database results contain inconsistent, missing, malformed, or unusual data, report the issue to the user.
                """
        )

while True:
    print("Escribe 'salir', 'quit' o 'exit' para salir")
    user_input = input(f"{MODEL} > ")
    if user_input.lower() in ['salir', 'quit', 'exit']:
        print("Adiosito")
        break
    print("Pensando...")
    agent_result = agent(user_input)
    print()
    print(f"context_size: {agent_result.context_size}")
    print(f"projected_context_size: {agent_result.projected_context_size}")