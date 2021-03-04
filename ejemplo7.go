package main

import "fmt"

func main() {
	// Operadores lógicos binarios
	var variable1 int8 // entero de 8 bits
	var variable2 int8 // entero de 8 bits
	var resultado int8 // entero de 8 bits

	variable1 = 12
	variable2 = 6

	fmt.Println("Variable1: ", variable1)
	fmt.Println("Variable2: ", variable2)

	resultado = variable1 | variable2
	fmt.Println("Resultado (OR): ", resultado)

	resultado = variable1 & variable2
	fmt.Println("Resultado (AND): ", resultado)

	resultado = variable1 ^ variable2
	fmt.Println("Resultado (XOR): ", resultado)

	// &^ es AND NOT, Ejemplo: A &^ B -> A AND (NOT B)
	resultado = variable1 &^ variable2
	fmt.Println("Resultado (AND NOT): ", resultado)

	// Tambien se pueden utilizar las asignaciones |=, &=, ^= y &^=
}
