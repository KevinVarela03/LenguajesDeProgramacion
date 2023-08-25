package main

import (
	"fmt"
	"sort"
)

// Definición de la estructura producto
type producto struct {
	nombre   string
	cantidad int
	precio   int
}

// Definición de la estructura listaProductos como una slice de producto
type listaProductos []producto

var lProductos listaProductos // Declaración de la lista de productos

const existenciaMinima int = 10 // Establecer existencia mínima para tomar decisiones

// Método para agregar un nuevo producto a la lista de productos
func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	var existe = lProductos.buscarProducto(nombre)

	if existe == -1 {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	} else {
		fmt.Println("El producto ya existe, pero se han modificado sus datos")
		var x int
		for x = 0; x < len(*l); x++ {
			if (*l)[x].nombre == nombre {
				(*l)[x].precio = precio
				(*l)[x].cantidad = cantidad
			}
		}
	}
	// Modificar el código para incrementar la cantidad y/o precio si el producto ya existe
}

// Método para buscar un producto por su nombre y retornar su índice en la lista
func (l *listaProductos) buscarProducto(nombre string) int {
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

// Método para vender productos y actualizar la cantidad en inventario
func (l *listaProductos) venderProducto(nombre string, cantidadVender int) {
	var prod = l.buscarProducto(nombre)
	var producto = l.buscarProducto(nombre)
	if cantidadVender > (*l)[producto].cantidad {
		fmt.Println("La cantidad es mayor que los productos en stock")
		prod = -1
	}
	if prod != -1 {
		(*l)[prod].cantidad = (*l)[prod].cantidad - cantidadVender
		if (*l)[prod].cantidad == 0 {
			var i int
			for i = 0; i < len(*l); i++ {
				if (*l)[i].nombre == (*l)[prod].nombre {
					fmt.Println("Se eliminó el producto de la lista: " + (*l)[i].nombre)
					lProductos = append(lProductos[:i], lProductos[i+1])
					break
				}
			}
		}
		// Modificar para eliminar el producto si no hay existencia y notificar la acción
		// Modificar para aceptar cantidad de productos a vender como parámetro
	}
}

// Función para llenar la lista de productos con datos iniciales
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 1, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

// Método para listar los productos con existencia mínima
func (l *listaProductos) listarProductosMínimos() listaProductos {
	var productosMinimos listaProductos

	for _, prod := range *l {
		if prod.cantidad <= existenciaMinima {
			productosMinimos = append(productosMinimos, prod)
		}
	}

	return productosMinimos
}

// Método para aumentar el inventario de productos mínimos hasta el límite
func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos listaProductos) {
	for _, prodMinimo := range listaMinimos {
		index := l.buscarProducto(prodMinimo.nombre)

		if index != -1 {
			diferencia := existenciaMinima - (*l)[index].cantidad
			if diferencia > 0 {
				(*l)[index].cantidad += diferencia
				fmt.Printf("Se aumentó el inventario de %s en %d unidades.\n", prodMinimo.nombre, diferencia)
			}
		}
	}
}

// Definición de tipo para ordenar productos por nombre
type ordenPorNombre listaProductos

// Implementación de los métodos de la interfaz sort.Interface para ordenar
func (o ordenPorNombre) Len() int           { return len(o) }
func (o ordenPorNombre) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o ordenPorNombre) Less(i, j int) bool { return o[i].nombre < o[j].nombre }

func main() {
	llenarDatos()
	fmt.Println("Lista de productos:")
	fmt.Println(lProductos)

	// Venta de productos
	lProductos.venderProducto("leche", 1)
	lProductos.venderProducto("frijoles", 5)

	fmt.Println("\nLista de productos después de ventas:")
	fmt.Println(lProductos)

	lProductos.agregarProducto("arroz", 20, 4000)
	fmt.Println("\nLista de productos después de agregar arroz:")
	fmt.Println(lProductos)

	lProductos.agregarProducto("Atun", 200, 2350)
	fmt.Println("\nLista de productos después de agregar Atun:")
	fmt.Println(lProductos)

	productosMinimos := lProductos.listarProductosMínimos()
	fmt.Println("\nLista de productos con existencia mínima:")
	fmt.Println(productosMinimos)

	// Aumentar inventario de productos mínimos
	lProductos.aumentarInventarioDeMinimos(productosMinimos)
	fmt.Println("\nLista de productos después de aumentar inventario:")
	fmt.Println(lProductos)

	// Ordenar por nombre
	sort.Sort(ordenPorNombre(lProductos))
	fmt.Println("\nLista de productos ordenada por nombre:")
	fmt.Println(lProductos)
}
