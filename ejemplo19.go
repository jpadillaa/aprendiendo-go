package main

import "fmt"

// Definición de un struct
type Persona struct {
	nombre    string
	apellidos string
	salario   int32
}

func main() {
	// Declaración de una instancia del struct Persona
	var empleado1 Persona

	// Se imprimer los valores por defecto del struct persona (Nota: Los zero values)
	fmt.Println("Empleado 1: ", empleado1)

	// Declaración e inicialización de una instancia del struct Persona
	empleado2 := Persona{"Jesse", "Padilla", 12500000}

	// Acceso a los campos de empleado 2
	fmt.Println("Empleado 2: ", empleado2)
	fmt.Println("Empleado 2 (Nombres): ", empleado2.nombre)
	fmt.Println("Empleado 2 (Apellidos): ", empleado2.apellidos)
	fmt.Println("Empleado 2 (Salarios): ", empleado2.salario)

	// Declaración e inicialización de una instancia del struct Persona
	// Inicialización nombrada
	empleado3 := Persona{nombre: "Obi Wan", apellidos: "Kenobi", salario: 66}

	fmt.Println("Empleado 3: ", empleado3)
	fmt.Println("Empleado 3 (Nombres): ", empleado3.nombre)
	fmt.Println("Empleado 3 (Apellidos): ", empleado3.apellidos)
	fmt.Println("Empleado 3 (Salarios): ", empleado3.salario)

	// Declaración e inicialización incompleta de una instancia del struct Persona
	// Campos inizalizados con su respetivo zero-value
	empleado4 := Persona{nombre: "Anakin"}
	fmt.Println("Empleado 4: ", empleado4)
}
