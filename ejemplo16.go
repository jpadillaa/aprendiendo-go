package main

import "fmt"

// variables global
var variable1 int = 999

func main() {
	// variables locales
	var variable1, variable2 int = 100, 50

	// Display the values of the variables
	fmt.Printf("\nVariable1 : %d", variable1)
	fmt.Printf("\nVariable2 : %d", variable2)

	imprimirGlobal()

	// aquí finaliza el alcance de la función principal y sus variables
}

func imprimirGlobal() {
	fmt.Printf("\nVariable1  (Global): %d", variable1)
}
