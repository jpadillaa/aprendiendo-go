package main

import "fmt"

func main() {
	// Estructura condicional - Switch
	var dia int8 = 4

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Ese no es un día valido de la semana!")
	}

	// Estructura condicional - Switch con expresión lógica asociada
	switch modulo := dia % 2; modulo {
	case 0:
		fmt.Println("Es un día par")
	default:
		fmt.Println("Es un día impar")
	}

	// Estructura condicional - Switch
	var mes int8 = 10
	switch {
	case mes <= 3:
		fmt.Println("Primer Trimestre")
	case mes > 3 && mes <= 6:
		fmt.Println("Segundo Trimestre")
	case mes > 6 && mes <= 9:
		fmt.Println("Tercer Trimestre")
	case mes > 9 && mes <= 12:
		fmt.Println("Cuarto Trimestre")
	default:
		fmt.Println("Este no es un mes valido")
	}

}
