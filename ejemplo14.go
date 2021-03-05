package main

import "fmt"

func multiplicar(numero1, numero2 int) {
	resultado := numero1 * numero2
	fmt.Println("Respuesta = ", resultado)
}

func saludar() {
	fmt.Println("Hola, Mundo!")
}

func main() {
	multiplicar(50, 5)

	/* defer - las declaraciones diferidas retrasan la ejecución de la función o método o un
	   método anónimo hasta que regresan las funciones cercanas. Diferir la función o los
	   argumentos de llamada de método se evalúan instantáneamente,
	   pero se ejecutan hasta que regresan las funciones cercanas.
	*/
	defer multiplicar(6, 4)
	saludar()
}
