// Doble backslash para hacer comentarios

// Contiene el paquete principal del programa
// Es el punto inicial para ejecutar el programa, por lo que es obligatorio escribirlo.
package main

// El paquete fmt package implementa las funciones de E/S.
// Similar a printf() and scanf() de C
import "fmt"

// La notación de las llaves de Go es parte de su sintaxis, si se cambia se genera un
// error de compilación.
func main() {
	// fmt.Println () función de biblioteca estándar para imprimir algo como una salida en
	// la pantalla. Si! la P es en mayuscula.
	fmt.Println("Hello World")

	// Para compilar y ejecutar el programa, en la terminal:
	// go build ejemplo1.go
	// ejemplo1.exe (Windows)
	// ./ejemplo1 (Linux)

	// Para ejecutar solamente (Compila automaticamente y almacena el ejecutable en un tmp)
	// go run ejemplo1.go
}
