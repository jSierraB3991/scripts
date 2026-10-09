import ollama

response = ollama.systemone(
    model='tev1:4b',
    state='Me cobraron el doble en marzo, y no me lo han devuelto, devuelvanmelo o voy a cancelar',
    questions={
        'is_go_it': {
            'type': 'noul',
            'instructions': 'El usuario está amenazando con irse?',
            'criteria': {
                'true': 'El usuario esta amezando con irse.',
                'false': 'El usuario no esta trando de cancelar.',
            },
        },
    },
)
print(response)
print(response.answers['is_go_it'].noul)