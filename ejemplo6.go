package main

import "fmt"

func main() {
	// Otros Operadores aritméticos
	var variable int16 // entero de 16 bits

	variable = 20

	fmt.Println("Variable: ", variable)

	variable += 10
	fmt.Println("Resultado (+=): ", variable)

	variable -= 15
	fmt.Println("Resultado (-=): ", variable)

	variable *= 10
	fmt.Println("Resultado (*=): ", variable)

	variable /= 5
	fmt.Println("Resultado (/=): ", variable)

	variable %= 5
	fmt.Println("Resultado (%=): ", variable)

}
