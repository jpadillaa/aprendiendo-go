package main

import "fmt"

func main() {
	/* Go proporciona un tipo de mapa que implementa una tabla hash, similar a los diccionarios de Python
	   La definición de un mapa o un diccionario usa la sintaxis map[Tipo de Dato de la Llave] Tipo de Dato del Valor

	   La función make asigna e inicializa el map y devuelve un valor de mapa y retorna un apuntador.
	   Si no se utiliza la función, es decir, solo declaramos el mapa, este tendra un valor nulo y no es
	   posible agregarle nuevos valores.

	*/
	// forma 1 de declarar e inicializar un mapa
	// var diccionario map[string]int
	// diccionario = make(map[string]int)

	// forma 2 de declarar e inicializar un mapa
	diccionario := make(map[string]int)

	// forma 1 de declarar e inicializar un mapa

	diccionario["Obi-Wan"] = 66
	diccionario["Anakin"] = 99
	diccionario["Mace"] = 99

	for llave, valor := range diccionario {
		fmt.Printf("\n%s - %d", llave, valor)
	}

	// consultar un llave valida
	valorSinRecorrido := diccionario["Obi-Wan"]
	fmt.Println("\nun valor: ", valorSinRecorrido)

	// consultar un llave invalida, siempre se recupera un zero-value pero no quiere decir
	// que el dato exista
	valorSinRecorrido = diccionario["Palpatine"]
	fmt.Println("\nun valor: ", valorSinRecorrido)

	// consultar un llave invalida, siempre se recupera un zero-value y validar que la
	// llave efectivamente no existe
	valorSinRecorrido, existe := diccionario["Palpatine"]
	fmt.Println("\nun valor: ", valorSinRecorrido)
	fmt.Println("existe esa llave?: ", existe)

	// eliminar una llave y un valor con delete
	delete(diccionario, "Mace")
	for llave, valor := range diccionario {
		fmt.Printf("\n%s - %d", llave, valor)
	}

}
