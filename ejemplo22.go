package main

import (
	"fmt"
)

func main() {
	var variable1 int
	var variable2 string
	var variable3 float32
	var variable4 bool

	/* La función fmt.Scanf () escanea los textos que se dan en la entrada estándar,
	lee desde allí y almacena los valores sucesivos separados por espacios en argumentos sucesivos
	según lo determinado por el tipo de dato.
	*/
	fmt.Scanf("%d", &variable1)
	fmt.Scanf("%s", &variable2)
	fmt.Scanf("%g", &variable3)
	fmt.Scanf("%t", &variable4)

	println("int: ", variable1)
	println("string: ", variable2)
	println("float: ", variable3)
	println("bool: ", variable4)
}
