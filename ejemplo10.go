package main

import "fmt"

// funcion1: Esta es una función que no recibe parámetro y no retorna valores
func funcion1() {
	fmt.Println("Una función sin parametros y sin retorno")
}

// funcion2: Esta es una función que recibe un parámetro de tipo string y no retorna valores
func funcion2(cadena string) {
	fmt.Println(cadena)
}

// funcion3: Esta es una función que recibe 3 parámetros de tipo entero y no retorna valores
func funcion3(valor1 int, valor2 int, valor3 int) {
	fmt.Printf("\nSuma = %d", valor1+valor2+valor3)
}

// funcion4: Es la misma funcion3 pero con la notación común en Go para especificar los tipos de los parámetros
func funcion4(valor1, valor2, valor3 int) {
	fmt.Printf("\nSuma = %d", valor1+valor2+valor3)
}

// funcion5: Es la misma funcion4 pero con la notación común en Go para especificar los tipos de los parámetros
func funcion5(valor1, valor2, valor3 int, cadena string) {
	fmt.Printf("\nResta = %d", valor1-valor2-valor3)
	fmt.Printf(cadena)
}

// funcion6: Una función que recibe dos parametros y retorna un valro de tipo enterio
func funcion6(valor1, valor2 int) int {
	return valor1 * valor2
}

// funcion7: Una función que recibe dos parametros y retorna un valro de tipo enterio
func funcion7(valor1, valor2 int) (int, int) {
	return valor1 * valor2, valor1 + valor2
}

func main() {
	// Invocando funciones que no retornan datos
	funcion1()
	funcion2("\nEjecutando una función que recibe un parámetro!")
	funcion3(100, 10, 1)
	funcion4(200, 20, 2)
	funcion5(300, 30, 3, "\nFunción con varios tipos de parámetros")

	// Invocando funciones que retornan datos
	r := funcion6(5, 10)
	fmt.Printf("\n%d", r)

	r, z := funcion7(5, 10)
	fmt.Printf("\n%d - %d", r, z)

	/* Invocando una función que retorna dos valores, pero con el operador _ descartamos
	   el dato que no es de interés.
	*/
	r, _ = funcion7(5, 10)
	fmt.Printf("\n%d", r)

}
