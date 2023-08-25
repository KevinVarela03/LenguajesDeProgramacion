package main

import (
	"fmt"
)

// Función para imprimir los espacios entre cada asterisco
func imprimirEspacios(espacios int) {
	for e := 0; e < espacios; e++ {
		fmt.Print("   ") // Imprime tres espacios para separar visualmente los asteriscos
	}
}

// Función para imprimir cada asterisco
func imprimirAsteriscos(asteriscos int) {
	for a := 0; a < asteriscos; a++ {
		fmt.Print("  *") // Imprime dos espacios antes de cada asterisco para centrarlo
	}
}

// Función que imprime la parte de arriba del dibujo
func imprimirParteArriba(cantAsteriscos int) {
	for i := 0; i < cantAsteriscos; i++ {
		espacios := (cantAsteriscos - i) - 1 // Cálculo de espacios necesarios
		asteriscos := (2 * i) + 1            // Cálculo de cantidad de asteriscos en esta línea

		imprimirEspacios(espacios)
		imprimirAsteriscos(asteriscos)
		fmt.Println() // Imprime una nueva línea al final de cada iteración
	}
}

// Función para imprimir la parte de abajo del dibujo
func imprimirParteAbajo(cantAsteriscos int) {
	for i := cantAsteriscos - 1; i > 0; i-- {
		espacios := cantAsteriscos - i // Cálculo de espacios necesarios
		asteriscos := (2 * i) - 1      // Cálculo de cantidad de asteriscos en esta línea

		imprimirEspacios(espacios)
		imprimirAsteriscos(asteriscos)
		fmt.Println() // Imprime una nueva línea al final de cada iteración
	}
}

func main() {

	var asteriscosLineaCentral int

	// Se solicita al usuario un número
	fmt.Println("\n\nIngrese el número de asteriscos en la línea del centro (impar y positivo): ")
	fmt.Scan(&asteriscosLineaCentral)

	// Condición por si el valor no es correcto
	if asteriscosLineaCentral%2 == 0 || asteriscosLineaCentral <= 0 {
		fmt.Println("\nEl valor debe ser un número impar y positivo")
		main() // Reinicia el proceso si el valor es inválido
	}

	// Se calcula la altura de la figura
	altura := (asteriscosLineaCentral + 1) / 2

	// Se imprime la parte superior e inferior de la figura
	imprimirParteArriba(altura)
	imprimirParteAbajo(altura)
}
