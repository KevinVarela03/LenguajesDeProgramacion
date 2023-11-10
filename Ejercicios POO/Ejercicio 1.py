# Clase que representa una hoja de documento gráfico
class Hoja:
    def __init__(self):
        self.objetos = []

    def agregarObjeto(self, objeto):
        self.objetos.append(objeto)

# Clase base para objetos representables
class ObjetoRepresentable:
    def __init__(self, nombre):
        self.nombre = nombre

# Subclase de ObjetoRepresentable para representar texto
class Texto(ObjetoRepresentable):
    def __init__(self, nombre, contenido):
        super().__init__(nombre)
        self.contenido = contenido

# Subclase de ObjetoRepresentable para representar formas geométricas
class Forma(ObjetoRepresentable):
    def __init__(self, nombre, color):
        super().__init__(nombre)
        self.color = color

# Subclases de Forma para representar un círculo y un cuadrado
class Circulo(Forma):
    def __init__(self, nombre, radio, color):
        super().__init__(nombre, color)
        self.radio = radio

class Cuadrado(Forma):
    def __init__(self, nombre, lado, color):
        super().__init__(nombre, color)
        self.lado = lado

# Subclase de ObjetoRepresentable para representar un grupo de objetos
class Grupo(ObjetoRepresentable):
    def __init__(self, nombre):
        super().__init__(nombre)
        self.objetos = []

    def agregarObjeto(self, objeto):
        self.objetos.append(objeto)

# Ejemplo de uso
hoja1 = Hoja()

texto1 = Texto("Texto 1", "Contenido del texto 1")
circulo1 = Circulo("Circulo 1", 10, "Rojo")
cuadrado1 = Cuadrado("Cuadrado 1", 5, "Azul")

grupo1 = Grupo("Grupo 1")
grupo1.agregarObjeto(texto1)
grupo1.agregarObjeto(circulo1)

hoja1.agregarObjeto(texto1)
hoja1.agregarObjeto(circulo1) 
hoja1.agregarObjeto(cuadrado1)
hoja1.agregarObjeto(grupo1)

# Recorremos los objetos de la hoja1 y sus grupos
for x in hoja1.objetos:
    if isinstance(x, Grupo):
        print(f"Grupo de la hoja: \n\t Nombre: {x.nombre}")
        for y in x.objetos:
            print(f"Objeto del grupo de la hoja: \n\t Nombre: {y.nombre}")
    else:
        print(f"Objeto de la hoja: \n\t Nombre: {x.nombre}")