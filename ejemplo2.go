package main

import "fmt"

func main() {
	// Ejemplo1: declaración de constantes
	const cpi1 float64 = 3.141592
	const cpi2 = 3.141592

	// Ejemplo2: Se crea una nueva variable llamada variable1 sin tipo explicito,
	// es una nueva variable (esto lo define el operador := ); es necesario asignar un valor
	variable1 := 12

	// Ejemplo3: Se crea una variable llamada variable2, de tipo entero y se inicializa con
	// un valor, es importante recalcar que al especificar el tipo el operador utilizado es =
	var variable2 int = 100

	// Ejemplo4: Se crea una variable llamada variable3, de tipo entero y sin valor inicial
	var variable3 int

	fmt.Println("PI1: ", cpi1)
	fmt.Println("PI2: ", cpi2)

	// Ejemplo5: Si no se utilizan las variables 1,2 y 3 al compilar o intentar ejecutar
	// go lanzara un error -> ejemplo2.go:12:2: variable1 declared but not used
	// En este ejemplo solo imprimo los valores en pantalla para poder ejecutar.
	fmt.Println("variable1: ", variable1)
	fmt.Println("variable2: ", variable2)
	fmt.Println("variable3: ", variable3)

}
