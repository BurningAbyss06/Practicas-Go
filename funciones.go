package main

import "fmt"

var funciones = map[string]func(int, int) int{
	"suma":           func(a int, b int) int { return a + b },
	"resta":          func(a int, b int) int { return a - b },
	"multiplicacion": func(a int, b int) int { return a * b },
	"division":       func(a int, b int) int { return a / b },
}

func imp(clave string, a int, b int) {
	operacion, exists := funciones[clave]
	if !exists {
		fmt.Println("Operacion invalida")
		return
	}

	fmt.Println("La", clave, "de", a, "y", b, "es", operacion(a, b))
}

func funcio() {
	imp("suma", 5, 6)
	imp("resta", 5, 6)
	imp("cddd", 5, 6)
	imp("multiplicacion", 5, 6)
	imp("division", 6, 2)
}
