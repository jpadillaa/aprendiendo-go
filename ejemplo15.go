package main

import "fmt"

func main() {
	// Declaración de Arrays
	var miArray [10]int
	miArrayFlotantes := [5]int{100, 200, 300, 400, 500}
	miArregloString := [5]string{"Fundamentos", "de", "programación", "con", "Go"}

	// Se imprime el zero value del array -> En este caso corresponde al del tipo int
	fmt.Println(miArray)
	fmt.Println(miArrayFlotantes)
	fmt.Println(miArregloString)

	for i := 0; i < 10; i++ {
		miArray[i] = i + 1
	}

	// Se imprime el array con cada elemento actualizado
	fmt.Println("\nmiArray modificado")
	fmt.Println(miArray)

	// la función len() retorna la longitud de un arreglo
	// la función cap() retorna la capacidad de un arreglo para almacenar elementos
	fmt.Println("len(): ", len(miArray))
	fmt.Println("cap(): ", cap(miArray))

	// arreglo definido por ellipsis, esto significa que el tamaño lo determinan los elementos en el arreglo
	miOtroArreglo := [...]int{111, 222, 333, 444, 555}
	fmt.Println("\nOriginal:", miOtroArreglo)

	// En Go los arreglos se copian por valor y no por referencia como pasa en Python
	miCopiaArreglo := miOtroArreglo
	fmt.Println("Copia:", miCopiaArreglo)

	miCopiaArreglo[0] = 999
	fmt.Println("Copia Modificado:", miCopiaArreglo)
	fmt.Println("Original para comparar:", miOtroArreglo)
	fmt.Println("Son iguales el arreglo original y la copia: ", miCopiaArreglo == miOtroArreglo)

	// Arreglo bi-dimensional o matriz de 2x2
	var matriz [2][2]int

	matriz[0][0] = 1
	matriz[0][1] = 0
	matriz[1][0] = 0
	matriz[1][1] = 1

	fmt.Println("\nMatriz")
	fmt.Println(matriz)

}
