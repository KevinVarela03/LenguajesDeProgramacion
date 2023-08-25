package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Estructura para representar el stock de un calzado
type stock struct {
	cantidad int     // Cantidad de calzados en stock
	Calzado  calzado // Información del calzado
}

// Lista de stocks
type listaStock []stock

var stocks listaStock

// Estructura para representar un calzado
type calzado struct {
	marca  string
	precio int
	talla  int
}

// Lista para guardar los calzados
type listaCalzados []calzado

// Variable para agregar un nuevo calzado
var calzados listaCalzados

// Función para agregar un nuevo calzado a la lista de calzados
func (c *listaCalzados) agregarCalzado(marca string, precio int, talla int) {
	// Verificar si el calzado ya existe
	var existe = c.buscarCalzado(marca, talla, precio)

	if existe == -1 {
		// Agregar un nuevo calzado si no existe
		*c = append(*c, calzado{marca: marca, precio: precio, talla: talla})
		fmt.Println("\nSe agrego un nuevo calzado. Cuenta con las siguientes caracteristicas:"+
			"\n\tMarca:", marca,
			"\n\tTalla:", talla,
			"\n\tPrecio:", precio)
	} else {
		// Aumentar la cantidad en stock si el calzado ya existe
		fmt.Println("\nEl calzado", marca, ", talla", talla, ", con precio", precio, "ya existe, pero se ha aumentado su cantidad en stock.")

		for i := range stocks {
			if stocks[i].Calzado.marca == marca && stocks[i].Calzado.talla == talla && stocks[i].Calzado.precio == precio {
				stocks[i].cantidad = stocks[i].cantidad + 1
				fmt.Println("\nAhora hay", stocks[i].cantidad, "calzados en stock.")
			}
		}
	}
}

// Función para llenar la lista de calzados con datos iniciales
func llenarDatos() {
	calzados.agregarCalzado("Nike", 75000, 42)
	calzados.agregarCalzado("Adidas", 80000, 34)
	calzados.agregarCalzado("Vans", 50000, 39)
	calzados.agregarCalzado("Hi-Tec", 90000, 40)
	calzados.agregarCalzado("Champion", 17000, 35)
	calzados.agregarCalzado("Timberland", 100000, 41)
	calzados.agregarCalzado("Caterpillar", 85000, 44)
	calzados.agregarCalzado("New Balance", 65000, 43)
	calzados.agregarCalzado("DC", 50000, 37)
	calzados.agregarCalzado("Crocs", 50000, 41)
	calzados.agregarCalzado("Puma", 45000, 38)
}

// Función para rellenar el stock basado en la lista de calzados
func rellenarStock() {
	for _, c := range calzados {
		if len(stocks) == 0 {
			numero := 5 //numeroRandom()
			nuevoStock := stock{
				cantidad: numero,
				Calzado:  c,
			}
			stocks = append(stocks, nuevoStock)
		} else {
			// Verificar si el calzado ya existe en el stock
			existeEnStock := false
			for p := 0; p < len(stocks); p++ {
				if stocks[p].Calzado.marca == c.marca && stocks[p].Calzado.talla == c.talla && stocks[p].Calzado.precio == c.precio {
					existeEnStock = true
					break
				}
			}
			// Si no existe, agregarlo con una cantidad aleatoria
			if !existeEnStock {
				numero := 5 //numeroRandom()
				nuevoStock := stock{
					cantidad: numero,
					Calzado:  c,
				}
				stocks = append(stocks, nuevoStock)
			}
		}
	}
}

// Función para generar un número aleatorio entre 0 y 50
func numeroRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(51)
}

// Función para buscar un calzado en la lista de calzados
func (c *listaCalzados) buscarCalzado(marca string, talla int, precio int) int {
	var resultado = -1
	for i := 0; i < len(*c); i++ {
		if (*c)[i].marca == marca && (*c)[i].talla == talla && (*c)[i].precio == precio {
			resultado = i
		}
	}
	return resultado
}

// Función para vender un calzado y actualizar el stock
func (s *listaStock) venderCalzado(marca string, precio int, talla int) {
	var existe = calzados.buscarCalzado(marca, talla, precio)

	if existe != -1 {
		for i := 0; i < len(*s); i++ {
			if (*s)[i].Calzado.marca == marca && (*s)[i].Calzado.talla == talla && (*s)[i].Calzado.precio == precio {
				if (*s)[i].cantidad > 0 {
					(*s)[i].cantidad = (*s)[i].cantidad - 1
					fmt.Println("\nSe han vendido un par de tenis de la marca", (*s)[i].Calzado.marca, ", talla", (*s)[i].Calzado.talla, "con precio de", (*s)[i].Calzado.precio, "colones")
					if (*s)[i].cantidad == 0 {
						fmt.Println("\nSe han agotado todos los productos en stock.")
					} else {
						fmt.Println("\nQuedan", (*s)[i].cantidad, "pares en stock.")
					}
					break
				} else {
					fmt.Println("\nYa no quedan calzados en stock, lo sentimos.")
				}
			}
		}
	} else {
		fmt.Println("\nEste producto no existe.")
	}
}

func main() {
	// Llenar la lista de calzados con datos iniciales
	llenarDatos()
	// Rellenar el stock basado en la lista de calzados
	rellenarStock()

	// Imprimir el stock actual
	fmt.Println("\nStock actual:\n", stocks)

	// Simular ventas y actualizaciones de stock
	stocks.venderCalzado("Nike", 75000, 42)
	stocks.venderCalzado("Nike", 75000, 42)
	stocks.venderCalzado("Nike", 75000, 42)
	stocks.venderCalzado("Nike", 75000, 42)
	stocks.venderCalzado("Nike", 75000, 42)

	// Agregar nuevos calzados y actualizar el stock
	calzados.agregarCalzado("Nike", 75000, 42)
	calzados.agregarCalzado("Adidas", 75000, 42)
	rellenarStock()

	// Imprimir el stock actual
	fmt.Println("\nStock actual:\n", stocks)

	// Simular más ventas y actualizaciones de stock
	stocks.venderCalzado("Caterpillar", 85000, 44)
	stocks.venderCalzado("Caterpillar", 85000, 44)
	stocks.venderCalzado("Caterpillar", 85000, 44)
	stocks.venderCalzado("Caterpillar", 85000, 44)
	stocks.venderCalzado("Caterpillar", 85000, 44)
	stocks.venderCalzado("Caterpillar", 85000, 44)

	// Imprimir el stock actual
	fmt.Println("\nStock actual:\n", stocks)
}
