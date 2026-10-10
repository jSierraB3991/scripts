
from fnmatch import fnmatch
from pathlib import Path

HOME = Path.home().resolve()

# Carpetas en las que nunca queremos entrar.
EXCLUDED_DIRS = {
    ".git",
    "node_modules",
    ".cache",
    "__pycache__",
    "venv",
    ".venv",
    ".vscode",
    ".ollama",
    ".local",
    ".lmstudio",
    "bin",
    "pkg",
    "build",
}

# Archivos exactos que queremos ignorar.
EXCLUDED_FILES = {
    "package-lock.json",
    "pnpm-lock.yaml",
}

# Patrones para excluir archivos.
EXCLUDED_FILE_PATTERNS = {
    "*.log",
    "*.tmp",
    "*.pyc",
    ".**",
}

def is_excluded_file(path: Path) -> bool:
    if path.name in EXCLUDED_FILES:
        return True

    return any(
        fnmatch(path.name, pattern)
        for pattern in EXCLUDED_FILE_PATTERNS
    )


def is_excluded_dir(path: Path) -> bool:
    return path.name in EXCLUDED_DIRS


def resolve_user_path(path: str) -> Path:
    requested = (HOME / path).resolve()

    if not requested.is_relative_to(HOME):
        raise ValueError(
            "La ruta debe estar dentro de la carpeta personal."
        )

    return requested


def list_files(path: str = ".") -> list[dict]:
    """Lista el contenido directo de una carpeta, sin entrar en subdirectorios."""

    directory = resolve_user_path(path)

    if not directory.is_dir():
        raise ValueError("La ruta indicada no es un directorio válido.")

    entries = []

    for item in sorted(
        directory.iterdir(),
        key=lambda p: p.name.lower(),
    ):
        try:
            is_dir = item.is_dir()

            if is_dir and is_excluded_dir(item):
                continue

            if not is_dir and is_excluded_file(item):
                continue

            entries.append({
                "name": item.name,
                "is_dir": is_dir,
                "size_bytes": (
                    item.stat().st_size if not is_dir else None
                ),
            })

        except OSError:
            continue

    return entries
