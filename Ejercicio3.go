package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Función para imprimir la lista rotada hacia la izquierda
func imprimirIzquierda(opcionLista string, cantRotaciones int, listaGenerica []string, listaNueva []string) {
	// Verificar si se debe trabajar con la lista genérica o la nueva
	if opcionLista == "g" {
		// Rotar la lista genérica hacia la izquierda
		var elemListaGenerica string
		for e := 0; e < cantRotaciones; e++ {
			elemListaGenerica = listaGenerica[0]
			listaGenerica = append(listaGenerica, elemListaGenerica)
			listaGenerica = append(listaGenerica[:0], listaGenerica[0+1:]...)
		}
		fmt.Println("\n", listaGenerica)
	} else if opcionLista == "n" {
		// Rotar la lista nueva hacia la izquierda
		var elemListaNueva string
		for e := 0; e < cantRotaciones; e++ {
			elemListaNueva = listaNueva[0]
			listaNueva = append(listaNueva, elemListaNueva)
			listaNueva = append(listaNueva[:0], listaNueva[0+1:]...)
		}
		fmt.Println("\n", listaNueva)
	} else {
		// Si la opción no es "g" ni "n", llamar a la función de rotaciones y dirección
		rotacionesYDireccion(opcionLista, listaGenerica, listaNueva)
	}
}

// Función para imprimir la lista rotada hacia la derecha
func imprimirDerecha(opcionLista string, cantRotaciones int, listaGenerica []string, listaNueva []string) {
	// Verificar si se debe trabajar con la lista genérica o la nueva
	if opcionLista == "g" {
		// Rotar la lista genérica hacia la derecha
		var elemListaGenerica string
		for e := 0; e < cantRotaciones; e++ {
			elemListaGenerica = listaGenerica[len(listaGenerica)-1]
			listaGenerica = append([]string{elemListaGenerica}, listaGenerica[0:]...)
			listaGenerica = listaGenerica[:len(listaGenerica)-1]
		}
		fmt.Println("\n", listaGenerica)
	} else if opcionLista == "n" {
		// Rotar la lista nueva hacia la derecha
		var elemListaNueva string
		for e := 0; e < cantRotaciones; e++ {
			elemListaNueva = listaNueva[len(listaNueva)-1]
			listaNueva = append([]string{elemListaNueva}, listaNueva[0:]...)
			listaNueva = listaNueva[:len(listaNueva)-1]
		}
		fmt.Println("\n", listaNueva)
	} else {
		// Si la opción no es "g" ni "n", llamar a la función de rotaciones y dirección
		rotacionesYDireccion(opcionLista, listaGenerica, listaNueva)
	}
}

// Función para preguntar por la cantidad de rotaciones y la dirección
func rotacionesYDireccion(opcionLista string, listaGenerica []string, listaNueva []string) {
	var opcionDireccion int
	var cantRotaciones int

	fmt.Print("\n¿Para cuál lado desea la rotación (izq = 0 y der = 1)?: ")
	fmt.Scan(&opcionDireccion)

	if opcionDireccion == 0 {
		fmt.Print("\n¿Cuál es la cantidad de posiciones que desea rotar?: ")
		fmt.Scan(&cantRotaciones)
		imprimirIzquierda(opcionLista, cantRotaciones, listaGenerica, listaNueva)
	} else if opcionDireccion == 1 {
		fmt.Print("\n¿Cuál es la cantidad de posiciones que desea rotar?: ")
		fmt.Scan(&cantRotaciones)
		imprimirDerecha(opcionLista, cantRotaciones, listaGenerica, listaNueva)
	} else {
		rotacionesYDireccion(opcionLista, listaGenerica, listaNueva)
	}
}

func main() {
	// Variables necesarias
	var opcionLista string
	listaGenerica := []string{}
	listaNueva := []string{}

	// Se pregunta al usuario si desea usar una lista existente o crear una nueva
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("\n\n¿Desea ingresar una secuencia nueva (n) o usar una generica (g)?: ")
	input, _ := scanner.ReadString('\n')
	opcionLista = strings.TrimSpace(input)

	if opcionLista == "n" {
		// Pedir al usuario que ingrese elementos para la lista nueva
		scanner2 := bufio.NewScanner(os.Stdin)
		fmt.Println("\nIngresa elementos que desea usar (presiona Enter en una línea vacía para finalizar):")
		for {
			scanner2.Scan()
			linea := scanner2.Text()
			if len(linea) == 0 {
				break
			}
			listaNueva = append(listaNueva, linea)
		}
		fmt.Println("\n", listaNueva)
	} else if opcionLista == "g" {
		// Usar lista genérica predefinida
		listaGenerica = []string{
			"a", "b", "c", "d", "e", "f", "g", "h",
		}
		fmt.Println("\n", listaGenerica)
	} else {
		// Si la opción no es "g" ni "n", reiniciar el proceso
		main()
	}

	// Llamar a la función para preguntar la dirección y la cantidad de rotaciones
	rotacionesYDireccion(opcionLista, listaGenerica, listaNueva)
}
