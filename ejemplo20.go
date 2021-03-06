package main

import "fmt"

//Definición de un Struct
type Persona struct {
	//Los mismos tipos de datos se pueden agrupar en una sola linea
	nombre, apellidos string
	salario           int32
}

func main() {
	// Declaración de una instancia del struct Persona
	var empleado1 Persona

	// Declaración de un apuntador (pointer) a una instancia de Persona
	var empleado2 *Persona

	empleado1 = Persona{"Jesse", "Padilla", 12500000}

	empleado2 = &empleado1

	fmt.Println("Empleado1:", empleado1)

	// imprime el apuntador y su contenido. Notese que el resultado contiene la misma información de empleado1,
	// sin embargo, Go coloca como prefijo el operador & para indicar que es un apuntador
	fmt.Println("Empleado2:", empleado2)

	// imprime el valor de a la estructura apuntada
	fmt.Println("Empleado2:", *empleado2)

	// (*empleado2).nombre es la sintaxis para acceder al valor nombre de la estructura a la que se apunta
	fmt.Println("Empleado 2 (Nombres): ", (*empleado2).nombre)
	fmt.Println("Empleado 2 (Apellidos): ", (*empleado2).apellidos)
	fmt.Println("Empleado 2 (Salarios): ", (*empleado2).salario)

	// (*empleado2).nombre utilizar esta notación no es mandatoria en Go. Go permita acceder a los campos
	// de una estructura a la que se apunta, como si fuera una instancia de una estructura declarada
	fmt.Println("Empleado 2 (Nombres): ", empleado2.nombre)
	fmt.Println("Empleado 2 (Apellidos): ", empleado2.apellidos)
	fmt.Println("Empleado 2 (Salarios): ", empleado2.salario)

}
