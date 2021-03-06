package main

import "fmt"

func main() {
	// Declaración de Arrays
	var miArray [10]int

	fmt.Println(miArray)
	respuesta := funcionRecorrerArray(miArray, 10)
	fmt.Println(respuesta)
	funcionModificarArray(&miArray, 10)
	fmt.Println(miArray)
	respuesta = funcionRecorrerArray(miArray, 10)
	fmt.Println(respuesta)
}

/* Esta función recibe el array por valor (una copia) y retorna la suma de todas las posiciones
   la variable size realmente sobra, pero el tamaño del arreglo en el parametro de la función NO
*/
func funcionRecorrerArray(miArray [10]int, size int) int {
	var r int

	for i := 0; i < size; i++ {
		r += miArray[i]
	}

	return r
}

/* Esta función recibe el array por referencia (un pointer) y actualiza el arreglo, no retorna nada
   dado que modifica al arreglo directamente.
   la variable size realmente sobra, pero el tamaño del arreglo en el parametro de la función NO
*/
func funcionModificarArray(miArray *[10]int, size int) {
	for i := 0; i < size; i++ {
		(*miArray)[i] = 5
	}
}
