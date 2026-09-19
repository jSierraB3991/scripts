from langchain_ollama import ChatOllama
import requests
from bs4 import BeautifulSoup
from pathlib import Path

def extract_text_of_url(url: str) -> str:
    print(f"Request to url: {url}")
    response = requests.get(url=url)
    soup = BeautifulSoup(response.content, 'html.parser')
    for tag in soup(["script", "style", "nav", "footer"]):
        tag.decompose()
    text = soup.get_text(separator='\n', strip=True)
    return text

def write_file(path: str, content: str) -> str:
    fiel = Path(path)
    try:
        fiel.write_text(content, encoding="utf-8")
        return f"Archivo escrito correctamente: {path}"
    except Exception as e:
        return f"Error escribiendfo en el archivo {path}"

def main():
    llm = ChatOllama(model="gemma4:12b")
    url = "https://edition.cnn.com"
    html_text = extract_text_of_url(url=url)
    print("search data of cnn in llm")
    response = llm.invoke(f"Este es el resultado extraido de la página de CNN. Analiza y descubre si hay alguna noticia relevante que pueda afectar a la bolsa {html_text}")
    print(response.content)
    print(write_file("./out_agent_02.txt",response.content))

if __name__  == "__main__":
    main()