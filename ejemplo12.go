package main

import "fmt"

func main() {

	fmt.Println("for tracional")

	// for tradicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("for tracional - decrementando")

	// for tradicional
	for i := 10; i > 10; i-- {
		fmt.Println(i)
	}

	fmt.Println("for tipo while")

	// for -> tipo while
	contador := 0
	for contador < 10 {
		fmt.Println(contador)
		contador++
	}

	fmt.Println("for forever")

	// for -> loop infinito (requiere una condición que lo rompa para evitar la ejecución infinita)
	contador = 0
	for {
		fmt.Println(contador)
		if contador > 10 {
			break
		}
		contador++

	}

}
