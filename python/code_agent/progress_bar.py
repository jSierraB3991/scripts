# -*- coding: utf-8 -*-
"""
Archivo con clase barra de carga
"""


class BarraCarga:
    """Clase para mostrar barras de carga simples"""
    
    def iniciar(self):
        """Empieza a cargar y muestra texto en color rojo"""
        
        print("\033[91mCargando\033[0m")
    
    def finalizar(self):
        """Finaliza la carga y muestra texto 'Listo'"""
        
        print("\033[92mListo\033[0m")