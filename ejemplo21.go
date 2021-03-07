package main

import "fmt"

// struct 1
type Estudiante struct {
	nombre, apellidos string
	codigo            int
}

// struct 2
type Curso struct {
	nombre, tema string
	lista        [3]Estudiante
}

func main() {
	// Se crea una instancia de Curso y se inicializa
	unCurso := Curso{
		nombre: "ISIS1304",
		tema:   "Fundamentos de Infraestructura Tecnólogica",
	}

	unCurso.lista[0] = Estudiante{"Obi-Wan", "Kenobi", 1}
	unCurso.lista[1] = Estudiante{"Anakin", "Skywalker", 2}
	unCurso.lista[2] = Estudiante{"Mace", "Windu", 3}

	fmt.Println("Detalles del curso")
	fmt.Println("Nombre: ", unCurso.nombre)
	fmt.Println("Tema: ", unCurso.tema)

	fmt.Println("\nLista de estudiantes")
	for i := 0; i < 3; i++ {
		fmt.Println("Nombre: ", unCurso.lista[i].nombre)
		fmt.Println("Apellido: ", unCurso.lista[i].apellidos)
		fmt.Println("Código: ", unCurso.lista[i].codigo)
	}

}
