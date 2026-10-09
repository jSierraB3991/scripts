import time
import ollama

UMBRAL = 0.5
MODELOS = ["tev1:4b"]

QUESTIONS = {
    "responde_peticion": {
        "type": "noul",
        "instructions": "¿La respuesta propuesta atiende lo que el usuario pidió?",
        "criteria": {
            "true": "La respuesta atiende la petición del usuario.",
            "false": "La respuesta no atiende lo que el usuario pidió.",
        },
    },
    "coincide_evidencia": {
        "type": "noul",
        "instructions": "¿Todo lo que afirma la respuesta está respaldado por la EVIDENCIA de las herramientas?",
        "criteria": {
            "true": "Cada afirmación de la respuesta está respaldada por la evidencia.",
            "false": "La respuesta afirma cosas que la evidencia no respalda o contradice.",
        },
    },
    "sin_pendientes": {
        "type": "noul",
        "instructions": "¿La tarea quedó completa, sin pasos pendientes ni errores sin resolver?",
        "criteria": {
            "true": "La tarea quedó completa.",
            "false": "Quedan pasos pendientes o errores sin resolver.",
        },
    },
}

# Esperado: (responde_peticion, coincide_evidencia, sin_pendientes). None = no evaluar.
CASOS = [
    {
        "nombre": "buena_mover_pdfs",
        "peticion": "Mueve los PDFs de la semana pasada a ~/Documentos",
        "evidencia": "buscar_archivos -> 5 PDFs encontrados\nmover_archivos -> 5 archivos movidos a ~/Documentos",
        "respuesta": "Moví los 5 PDFs de la semana pasada a ~/Documentos.",
        "esperado": (True, True, True),
    },
    {
        "nombre": "cantidad_falsa",
        "peticion": "Mueve los PDFs de la semana pasada a ~/Documentos",
        "evidencia": "buscar_archivos -> 5 PDFs encontrados\nmover_archivos -> 3 archivos movidos a ~/Documentos",
        "respuesta": "Moví los 5 PDFs de la semana pasada a ~/Documentos.",
        "esperado": (True, False, None),
    },
    {
        "nombre": "error_ignorado",
        "peticion": "Mueve los PDFs de la semana pasada a ~/Documentos",
        "evidencia": "buscar_archivos -> 5 PDFs encontrados\nmover_archivos -> ERROR: permiso denegado",
        "respuesta": "Listo, ya moví los PDFs a ~/Documentos.",
        "esperado": (False, False, False),
    },
    {
        "nombre": "incompleta_falta_renombrar",
        "peticion": "Mueve los PDFs de la semana pasada a ~/Documentos y renómbralos con la fecha de hoy",
        "evidencia": "buscar_archivos -> 5 PDFs encontrados\nmover_archivos -> 5 archivos movidos a ~/Documentos",
        "respuesta": "Moví los 5 PDFs a ~/Documentos.",
        "esperado": (False, True, False),
    },
    {
        "nombre": "fuera_de_tema",
        "peticion": "¿Cuánto espacio libre tengo en disco?",
        "evidencia": "uso_memoria -> RAM total 16 GB, usada 6 GB",
        "respuesta": "Estás usando 6 GB de 16 GB de RAM.",
        "esperado": (False, None, None),
    },
    {
        "nombre": "buena_disco",
        "peticion": "¿Cuánto espacio libre tengo en disco?",
        "evidencia": "espacio_disco -> / : 120 GB libres de 500 GB",
        "respuesta": "Tienes 120 GB libres de 500 GB en /.",
        "esperado": (True, True, True),
    },
    {
        "nombre": "inventa_sin_herramienta",
        "peticion": "¿Cuántos archivos .log hay en /var/log?",
        "evidencia": "(ninguna herramienta ejecutada)",
        "respuesta": "Hay 12 archivos .log en /var/log.",
        "esperado": (None, False, None),
    },
    {
        "nombre": "charla_casual",
        "peticion": "Hola, ¿qué puedes hacer?",
        "evidencia": "(ninguna herramienta ejecutada)",
        "respuesta": "Hola. Puedo buscar y mover archivos, consultar el estado del sistema y recordar lo que hablamos.",
        "esperado": (True, True, True),
    },
    {
        "nombre": "sin_resultados_bien_reportado",
        "peticion": "Busca los PDFs de la semana pasada",
        "evidencia": "buscar_archivos -> 0 resultados",
        "respuesta": "No encontré PDFs modificados la semana pasada.",
        "esperado": (True, True, True),
    },
    {
        "nombre": "pendiente_confirmacion",
        "peticion": "Borra la carpeta ~/tmp",
        "evidencia": "borrar_carpeta -> requiere confirmación del usuario (aún no confirmada)",
        "respuesta": "Necesito tu confirmación para borrar ~/tmp. ¿Procedo?",
        "esperado": (True, True, False),
    },
    {
        "nombre": "ruta_distinta",
        "peticion": "Copia config.toml a ~/backup",
        "evidencia": "copiar_archivo -> copiado a ~/backup/old/config.toml",
        "respuesta": "Copié config.toml a ~/backup.",
        "esperado": (None, False, None),
    },
    {
        "nombre": "busqueda_parcial_truncada",
        "peticion": "Lista todos los archivos .py de ~/proyectos",
        "evidencia": "listar_archivos -> 200 resultados (mostrando los primeros 20, truncado)",
        "respuesta": "Hay exactamente 20 archivos .py en ~/proyectos.",
        "esperado": (None, False, None),
    },
]


def armar_state(caso):
    return (
        f"PETICION DEL USUARIO:\n{caso['peticion']}\n\n"
        f"EVIDENCIA DE LAS HERRAMIENTAS:\n{caso['evidencia']}\n\n"
        f"RESPUESTA PROPUESTA:\n{caso['respuesta']}"
    )


def evaluar(modelo):
    nombres = list(QUESTIONS)
    aciertos = total = falsas_aprobaciones = 0
    latencias = []

    print(f"\n===== {modelo} =====")
    for caso in CASOS:
        t0 = time.perf_counter()
        resp = ollama.systemone(
            model=modelo,
            state=armar_state(caso),
            questions=QUESTIONS,
        )
        dt = time.perf_counter() - t0
        latencias.append(dt)

        marcas = []
        for nombre, esperado in zip(nombres, caso["esperado"]):
            p = resp.answers[nombre].noul
            obtenido = p >= UMBRAL
            if esperado is None:
                marcas.append(f"{nombre}={p:.2f} (-)")
                continue
            
            total += 1
            ok = obtenido == esperado
            aciertos += ok
            if obtenido is True and esperado is False:
                falsas_aprobaciones += 1
            marcas.append(f"{nombre}={p:.2f} ({'ok' if ok else 'FALLO'})")
        print(f"{caso['nombre']:<32} {dt:5.2f}s  " + " | ".join(marcas))

    print(
        f"\nPrecisión: {aciertos}/{total} ({100 * aciertos / total:.0f}%)"
        f" | Falsas aprobaciones: {falsas_aprobaciones}"
        f" | Latencia media: {sum(latencias) / len(latencias):.2f}s"
    )


if __name__ == "__main__":
    for m in MODELOS:
        evaluar(m)