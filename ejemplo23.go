package main

import "fmt"

func main() {
	/* Un slice es una secuencia de longitud variable que almacena elementos de un tipo similar,
	   no se le permite almacenar diferentes tipos de elementos en el mismo segmento.
	   Es como una array que tiene un valor de índice y una longitud, pero de tamaño variable y no fijo como un array.
	*/
	var miSlice []int

	miSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(miSlice)

	// slicing  (se lista del elemento 4 al 6, no se incluye el 3 - similar a python)
	miSlice2 := miSlice[3:6]
	fmt.Println(miSlice2)

	// slicing (se lista desde el elemento 7 hasta el fin)
	miSlice2 = miSlice[6:]
	fmt.Println(miSlice2)

	// slicing
	miSlice2 = miSlice[:5]
	fmt.Println(miSlice2)

	// los arrays son inmutables por lo que no se les puede agregar datos, al slicing si con la función append
	miSlice = append(miSlice, 16, 17, 18, 19, 20)
	fmt.Println(miSlice)

	// Con append puedo concatenar dos slides, los ... son parte de la sintaxis
	otroSlice := []int{21, 22, 23, 24, 25}
	miSlice = append(miSlice, otroSlice...)
	fmt.Println(miSlice)

	// eliminando elementos con append
	miSlice = append(miSlice[:10], miSlice[20:]...)
	fmt.Println(miSlice)

	// recorriendo slides con range, se puede utilizar _ para descartar el índice o el valor si alguno no es de interés.
	for indice, valor := range miSlice {
		fmt.Println("indice: ", indice)
		fmt.Println("valor: ", valor)
	}

}
