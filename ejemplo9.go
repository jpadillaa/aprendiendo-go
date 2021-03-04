package main

import "fmt"

func main() {
	var cadena1 string
	var cadena2 string
	var variable1 int8
	var variable2 int8

	cadena1 = "Hola,"
	cadena2 = "Mundo!"

	variable1 = 64   // #64 en decimal
	variable2 = 0x3F // #63 representado en hexadecimal

	fmt.Println(cadena1, cadena2)

	// Printf es una  función más versatil, y funciona similar a su par en C
	fmt.Printf("\n%s %s", cadena1, cadena2)

	// %d opera igual que en C, se utiliza para imprimir un int
	fmt.Printf("\n%d + %d = %d", variable1, variable2, variable1+variable2)

	/* %s opera igual que en C, se utiliza para imprimir una cadena
	   %v se utiliza cuando no se tiene claro el tipo de dato que se desea imprimir
	   por cierto, esto es un comentario de múltiples líneas
	*/
	fmt.Printf("\n%s %v + %v = %v", cadena1, variable1, variable2, variable1+variable2)

	/* Sprinft se utiliza para generar una cadena a partir de un especificador de formato
	retorna una cadena de texto */
	// mensaje es una variable que no se declaro previamente
	mensaje := fmt.Sprintf("\n%s %v + %v = %v", cadena1, variable1, variable2, variable1+variable2)
	fmt.Println(mensaje)
}
