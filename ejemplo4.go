package main

import "fmt"

func main() {
	// Profundizando en los tipos de datos
	// tipo: int  (con signo)
	var variable1 int   // enteros de 64 bits
	var variable2 int8  // enteros de 8 bits
	var variable3 int16 // enteros de 16 bits
	var variable4 int32 // enteros de 32 bits
	var variable5 int64 // enteros de 34 bits

	// tipo: int  (sin signo)
	var uVariable1 uint   // enteros sin signo de 64 bits
	var uVariable2 uint8  // enteros sin signo de 8 bits
	var uVariable3 uint16 // enteros sin signo de 16 bits
	var uVariable4 uint32 // enteros sin signo de 32 bits
	var uVariable5 uint64 // enteros sin signo de 34 bits

	// tipo: float
	var fVariable1 float32 // decimales de 32 bits
	var fVariable2 float64 // decimales de 64 bits

	// tipo: string
	var sVariable1 string                 // string vacio
	var sVariable2 string = ""            // string vacio, siempre ""
	var sVariable3 string = "Una cadena!" // string con una cadena

	// tipo: string
	var bVariable1 bool         // bool con su valor defecto
	var bVariable2 bool = false // bool con su valor falso
	var bVariable3 bool = true  // bool con su valor verdadero

	// tipo: complex
	var cVariable1 complex64          // Complex64 con su valor defecto float32 real, float32 img
	var cVariable2 complex128         // Complex128 con su valor defecto float64 real, float64 img
	var cVariable3 complex64 = 1 + 1i // Complex64 con un ejemplo

	// imprimiendo enteros con signo
	fmt.Println("Entero (int): ", variable1)
	fmt.Println("Entero (int8): ", variable2)
	fmt.Println("Entero (int16): ", variable3)
	fmt.Println("Entero (int32): ", variable4)
	fmt.Println("Entero (int64): ", variable5)

	// imprimiendo enteros sin signo
	fmt.Println("Entero (uint): ", uVariable1)
	fmt.Println("Entero (uint8): ", uVariable2)
	fmt.Println("Entero (uint16): ", uVariable3)
	fmt.Println("Entero (uint32): ", uVariable4)
	fmt.Println("Entero (uint64): ", uVariable5)

	// imprimiendo decimales
	fmt.Println("Decimal (float32): ", fVariable1)
	fmt.Println("Decimal (float64): ", fVariable2)

	// imprimiendo strings
	fmt.Println("string (defecto): ", sVariable1)
	fmt.Println("string (vacio): ", sVariable2)
	fmt.Println("string (cadena): ", sVariable3)

	// imprimiendo booleans
	fmt.Println("bool (defecto): ", bVariable1)
	fmt.Println("bool (false): ", bVariable2)
	fmt.Println("bool (true): ", bVariable3)

	// imprimiendo Complex
	fmt.Println("complex64 (defecto): ", cVariable1)
	fmt.Println("complex128 (defecto): ", cVariable2)
	fmt.Println("complex64 (Ejemplo): ", cVariable3)
}
