package main

import "fmt"

func main() {
	/* Los operadores de comparación son los mismos de cualquier otro lenguaje

	valor1 == valor2: TRUE si valor1 y valor2 son iguales.
	valor1 != valor2: TRUE si valor1 es diferente de valor2.
	valor1 < valor2: TRUE si valor1 es menor que valor2
	valor1 > valor2: TRUE si valor1 es mayor que valor2
	valor1 >= valor2: TRUE si valor1 es igual o mayor que valor2
	valor1 <= valor2: TRUE si valor1 es menor o igual que valor2.

	Se pueden componer expresiones más complejas con && -> AND, || -> OR y !(var_bool) -> NOT
	*/

	var variable1 int8 = 8

	/* Nótese que el else va en la misma línea de la llave que cierra el bloque if
	   -> no hacerlo de esta forma genera error.

	   Aplican las mismas reglas de llaves de las funciones
	*/
	if variable1%2 == 0 {
		fmt.Println("El numero es par!")
	} else {
		fmt.Println("El numero es impar!")
	}

	// Al igual que en otros lenguajes el bloque else no es mandatorio
	variable1 = 9
	resultado := "El numero es impar"

	if variable1%2 == 0 {
		resultado = "El numero es par"
	}
	fmt.Println(resultado)

	// Sentencia if - else if - else
	variable1 = 15
	if variable1%2 == 0 {
		fmt.Println("El numero es par!")
	} else if variable1%3 == 0 {
		fmt.Println("El numero es divisible por 3!")
	} else {
		fmt.Println("El numero no es par y no es divisible por 3!")
	}
}
