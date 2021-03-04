package main

import "fmt"

func main() {
	var variable1 int8 // entero de 8 bits
	var variable2 int8 // entero de 8 bits
	var variable3 int8 // entero de 8 bits

	variable1 = 64   // #64 en decimal
	variable2 = 0x40 // #64 representado en hexadecimal
	variable3 = 0100 // #64 representado en octal

	fmt.Println("Variable1: ", variable1)
	fmt.Println("Variable2: ", variable2)
	fmt.Println("Variable3: ", variable3)
}
