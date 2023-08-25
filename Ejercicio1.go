package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Funcion para hacer el conteo
func conteo(texto string) {

	// Variables para el recorrido
	var c int // Contador de caracteres
	var p int // Contador de patrones de palabras
	var l int // Contador de líneas

	// Variables para almacenar el total de datos
	var caracteres int
	var palabras int
	var lineas = 1 // Inicializar en 1 ya que hay al menos una línea

	// Se calcula la cantidad de caracteres sin espacios
	for c = 0; c < len(texto); c++ {
		if texto[c] == 32 || texto[c] == 10 {
			// Ignorar espacios y nuevas líneas
		} else {
			caracteres++
		}
	}

	// Se calcula la cantidad de palabras
	for p = 0; p < len(texto); p++ {
		for p+1 < len(texto) {
			// Verificar si el carácter actual y el siguiente forman un delimitador de palabra
			if texto[p] >= 32 && texto[p] < 65 || texto[p] > 91 && texto[p] < 96 || texto[p] > 123 && texto[p] < 126 || texto[p+1] == 10 {
				// Verificar si el carácter anterior era una letra
				if texto[p-1] >= 65 && texto[p-1] <= 90 || texto[p-1] >= 97 && texto[p-1] <= 122 {
					palabras++
				}
			}
			p++
		}
		palabras++
	}

	// Se calcula la cantidad de lineas de texto
	for l = 0; l < len(texto); l++ {
		if texto[l] == 10 {
			lineas++
		}
	}

	// Se imprimen los resultados
	fmt.Print("\nLa cantidad de caracteres del texto '"+texto+"' son ", caracteres)
	fmt.Println("\n\nLa cantidad de palabras del texto '"+texto+"' son ", palabras)
	fmt.Println("\nLa cantidad de lineas de texto del texto '"+texto+"' son ", lineas)
}

func main() {

	// Se pregunta al usuario
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("\n\n¿Desea ingresar un texto nuevo (n) o usar uno generico (g)?: ")
	input, _ := scanner.ReadString('\n')
	input = strings.TrimSpace(input)

	// Se verifica si el usuario elige ingresar un texto nuevo o usar uno existente
	if input == "n" {
		scanner2 := bufio.NewScanner(os.Stdin)
		fmt.Println("\nIngresa el texto (presiona Enter en una línea vacía para finalizar):")

		var lines []string
		for {
			scanner2.Scan()
			line := scanner2.Text()
			if len(line) == 0 {
				break
			}
			lines = append(lines, line)
		}

		// Unir las líneas en un solo texto
		inputText := strings.Join(lines, "\n")

		if len(lines) == 0 {
			fmt.Print("\nLa cantidad de caracteres del texto son 0")
			fmt.Println("\n\nLa cantidad de palabras del texto son 0")
			fmt.Println("\nLa cantidad de lineas de texto del texto son 0")
		} else {
			conteo(inputText)
		}
	} else if input == "g" {
		// Texto genérico predefinido
		conteo("sigla en ingles de\n American Standard Code for Information Interchange\n (Codigo Estadounidense Estandar para el Intercambio de Informacion)")
	} else {
		// Reiniciar el proceso si no se elige "n" ni "g"
		main()
	}
}
