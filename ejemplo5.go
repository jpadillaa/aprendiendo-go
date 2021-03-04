package main

import "fmt"

func main() {
	// Operadores aritméticos y de corrimiento binario
	var variable1 int8 // entero de 8 bits
	var variable2 int8 // entero de 8 bits
	var resultado int8 // entero de 8 bits

	variable1 = 15
	variable2 = 2

	fmt.Println("Variable1: ", variable1)
	fmt.Println("Variable2: ", variable2)

	resultado = variable1 + variable2
	fmt.Println("Resultado (+): ", resultado)

	resultado = variable1 - variable2
	fmt.Println("Resultado (-): ", resultado)

	resultado = variable1 * variable2
	fmt.Println("Resultado (*): ", resultado)

	resultado = variable1 / variable2
	fmt.Println("Resultado (/): ", resultado)

	resultado = variable1 % variable2
	fmt.Println("Resultado (%): ", resultado)

	resultado = variable1 >> variable2
	fmt.Println("Resultado (>>): ", resultado)

	resultado = variable1 << variable2
	fmt.Println("Resultado (<<): ", resultado)

	variable1++
	fmt.Println("Resultado (++): ", variable1)

	variable1--
	fmt.Println("Resultado (++): ", variable1)

}
