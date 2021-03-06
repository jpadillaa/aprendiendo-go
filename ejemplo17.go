package main

import "fmt"

func main() {
	var variable1 int32 = 999
	var variable2 int32 = 100
	var variable3 float32
	// casting: conversión de tipo
	var variable4 int8 = int8(variable1)

	// conversión de tipo explicita
	variable3 = float32(variable1) / float32(variable2)

	fmt.Printf("Resultado = %f\n", variable3)
	fmt.Printf("Variable4 (Truncada) = %d\n", variable4)

	// conversión de float a int y uint
	var1, var2, var3 := 117.5, 78.34, -50.85
	fmt.Printf("%d %d %d\n", int(var1), int8(var2), uint8(var3))
}
