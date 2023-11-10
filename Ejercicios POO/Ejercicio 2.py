# Clase para representar a un Socio
class Socio:
    def __init__(self, nombre, id, direccion):
        self.nombre = nombre
        self.id = id
        self.direccion = direccion
        self.libros = [] # Lista de libros prestados al socio

# Clase para representar un Libro
class Libro:
    def __init__(self, titulo, autor, codigo, disponibilidad = True, localizacion = None):
        self.nombre = titulo
        self.autor = autor
        self.codigo = codigo
        self.disponibilidad = disponibilidad # Indica si el libro está disponible
        self.localizacion = localizacion # Socio que tiene el libro prestado (si no está disponible, será None)

# Clase para representar un Prestamo
class Prestamo: 
    def __init__(self,  socio, libro, fechaPrestamo):
        self.socio = socio
        self.libro = libro
        self.fechaPrestamo = fechaPrestamo

# Clase para representar una Biblioteca
class Biblioteca:
    def __init__(self):
        self.socios = []  # Lista de socios registrados
        self.libros = []  # Lista de libros disponibles en la biblioteca
        self.prestamos = []  # Lista de préstamos realizados
        
    def agregarSocio(self, socio):
        self.socios.append(socio) # Agregar un socio a la lista de socios
    
    def mostrarSocios(self):
        print("\n---LISTA DE SOCIOS--- ")
        for x in self.socios:
            print(f"\n\t Nombre: {x.nombre} \n\t ID: {x.id} \n\t Direccion: {x.direccion} \n\t Libros:")
            for y in x.libros:
                print(f"\t\t Nombre: {y.nombre} \n\t\t Autor: {y.autor} \n\t\t Codigo: {y.codigo} \n")

    def agregarLibro(self, libro):
        self.libros.append(libro) # Agregar un libro a la lista de libros disponibles

    def agregarPrestamo(self, socio, libro, fechaPrestamo):
        if libro.disponibilidad == False:
            print(f"\n!!!EL LIBRO '{libro.nombre}' NO ESTA DISPONIBLE!!!\n")
        else:
            prestamo = Prestamo(socio, libro, fechaPrestamo)  # Crear un registro de préstamo
            self.prestamos.append(prestamo)  # Agregar el préstamo a la lista de préstamos
            libro.disponibilidad = False  # Marcar el libro como no disponible
            libro.localizacion = socio  # Establecer al socio como la localización del libro
            socio.libros.append(libro)  # Agregar el libro prestado a la lista de libros del socio

    def devolverLibro(self, libro):
        libro.disponibilidad = True  # Marcar el libro como disponible
        for x in libro.localizacion.libros:
            if x.codigo == libro.codigo:
                libro.localizacion.libros.remove(x)  # Quitar el libro de la lista de libros del socio
                libro.localizacion = None  # Establecer la localización del libro como None
                for y in self.prestamos:
                    if y.libro.codigo == libro.codigo:
                        self.prestamos.remove(y)  # Eliminar el registro de préstamo
                        print(f"\n!!!EL LIBRO '{libro.nombre}' HA SIDO DEVUELTO!!!\n")
                break

    def mostrarLibros(self):
        print('\n---LISTA DE LIBROS--- ')
        for x in self.libros:
            print(f"\n\t Nombre: {x.nombre} \n\t Autor: {x.autor} \n\t Codigo: {x.codigo} \n\t Disponibilidad: {'Disponible' if x.disponibilidad else 'No disponible'} \n\t Localizacion: {x.localizacion.nombre if x.localizacion else 'Biblioteca'} \n")

    def mostrarPrestamos(self):
        print('\n---LISTA DE PRESTAMOS--- ')
        for x in self.prestamos:
            print(f"\n\t Socio: {x.socio.nombre} \n\t Libro: {x.libro.nombre} \n\t Fecha de prestamo: {x.fechaPrestamo} \n")

    def sociosConMasDe3Libros(self):
        socios = list(filter(lambda x: len(x.libros) > 3, self.socios))
        print('\n---LISTA DE SOCIOS CON MAS DE 3 LIBROS--- ')
        for x in socios:
            print(f"\n\t Nombre: {x.nombre} \n\t ID: {x.id} \n\t Direccion: {x.direccion} \n\t Cantidad de Libros: {len(x.libros)} \n\t Libros:")
            for y in x.libros:
                print(f"\t\t Nombre: {y.nombre} \n\t\t Autor: {y.autor} \n\t\t Codigo: {y.codigo} \n")

#Crear una instancia de la Biblioteca
biblioteca = Biblioteca()

#Crear socios
socio1 = Socio("Juan", 1, "Calle 1")
socio2 = Socio("Pedro", 2, "Calle 2")
socio3 = Socio("Maria", 3, "Calle 3")

#Agregar socios a la biblioteca
biblioteca.agregarSocio(socio1)
biblioteca.agregarSocio(socio2)
biblioteca.agregarSocio(socio3)

#Crear libros
libro1 = Libro("Libro 1", "Autor 1", 1)
libro2 = Libro("Libro 2", "Autor 2", 2)
libro3 = Libro("Libro 3", "Autor 3", 3)
libro4 = Libro("Libro 4", "Autor 4", 4)
libro5 = Libro("Libro 5", "Autor 5", 5)
libro6 = Libro("Libro 6", "Autor 6", 6)
libro7 = Libro("Libro 7", "Autor 7", 7)
libro8 = Libro("Libro 8", "Autor 8", 8)
libro9 = Libro("Libro 9", "Autor 9", 9)

#Agregar libros
biblioteca.agregarLibro(libro1)
biblioteca.agregarLibro(libro2)
biblioteca.agregarLibro(libro3)
biblioteca.agregarLibro(libro4)
biblioteca.agregarLibro(libro5)
biblioteca.agregarLibro(libro6)
biblioteca.agregarLibro(libro7)
biblioteca.agregarLibro(libro8)
biblioteca.agregarLibro(libro9)

#Agrergar préstamos de libros
biblioteca.agregarPrestamo(socio1, libro1, "01/01/2021")
biblioteca.agregarPrestamo(socio2, libro2, "01/01/2021")
biblioteca.agregarPrestamo(socio3, libro3, "01/01/2021")
biblioteca.agregarPrestamo(socio1, libro4, "01/01/2021")
biblioteca.agregarPrestamo(socio2, libro5, "01/01/2021")
biblioteca.agregarPrestamo(socio3, libro6, "01/01/2021")
biblioteca.agregarPrestamo(socio1, libro7, "01/01/2021")
biblioteca.agregarPrestamo(socio2, libro8, "01/01/2021")
biblioteca.agregarPrestamo(socio2, libro9, "01/01/2021")

#Libro ya en uso
biblioteca.agregarPrestamo(socio3, libro1, "01/01/2021")

#Devolver libros
biblioteca.devolverLibro(libro7)
biblioteca.devolverLibro(libro3)
biblioteca.devolverLibro(libro1)

#Mostrar socios, libros, préstamos y socios con más de 3 libros
biblioteca.mostrarSocios()
biblioteca.mostrarLibros()
biblioteca.mostrarPrestamos()
biblioteca.sociosConMasDe3Libros()