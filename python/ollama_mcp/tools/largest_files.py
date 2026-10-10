
from pathlib import Path

from .file_listing import (
    EXCLUDED_DIRS,
    EXCLUDED_FILES,
    EXCLUDED_FILE_PATTERNS,
    HOME,
    is_excluded_file,
    resolve_user_path,
)


def find_largest_files(
    path: str = ".",
    limit: int = 5,
    recursive: bool = True,
) -> list[dict]:
    """
    Encuentra los archivos más grandes dentro de una carpeta.

    Args:
        path: Carpeta inicial, relativa al directorio personal.
        limit: Número máximo de archivos que se devolverán.
        recursive: Si True, busca también en subdirectorios.
    """

    if not 1 <= limit <= 100:
        raise ValueError("limit debe estar entre 1 y 100.")

    root = resolve_user_path(path)

    if not root.is_dir():
        raise ValueError("La ruta indicada no es un directorio válido.")

    largest = []

    def scan(directory: Path) -> None:
        try:
            with __import__("os").scandir(directory) as entries:
                for entry in entries:
                    item = Path(entry.path)

                    try:
                        if entry.is_symlink():
                            continue

                        if entry.is_dir(follow_symlinks=False):
                            if recursive and item.name not in EXCLUDED_DIRS:
                                scan(item)

                            continue

                        if not entry.is_file(follow_symlinks=False):
                            continue

                        if is_excluded_file(item):
                            continue

                        size = entry.stat(follow_symlinks=False).st_size

                        largest.append({
                            "name": item.name,
                            "path": str(item),
                            "size_bytes": size,
                        })

                    except OSError:
                        continue

        except OSError:
            return

    scan(root)

    largest.sort(
        key=lambda file: file["size_bytes"],
        reverse=True,
    )

    return largest[:limit]
