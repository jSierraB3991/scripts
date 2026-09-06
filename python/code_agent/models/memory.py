class Memory:
    """Clase que representa una memoria guardada en la base de datos."""
    
    def __init__(self, id: int, key: str, content: str, role_agent: str, description: str, created_at: str):
        self.id = id
        self.key = key
        self.content = content
        self.created_at = created_at
        self.role_agent = role_agent
        self.description = description
    
    def to_dict(self) -> dict:
        """Convierte el objeto Memory a un diccionario."""
        return {
            "id": self.id,
            "key": self.key,
            "content": self.content,
            "created_at": self.created_at,
            "role_agent": self.role_agent,
        }
    def to_model(self) -> str:
        return f"{self.content} es el valor para: {self.description}, no tienes que guardarlo en base de datos, ya que este dato viene de allí, si quieres modificarlo puedes hacer con la 'key' {self.key}"