package main

import "fmt"

func puntero() {

	var x, y *int
	entero := 5
	x = &entero
	y = &entero

	fmt.Println("X apunta a:", x)
	fmt.Println("y apunta a:", y)
	fmt.Println("El valor de x es:", *x)
	fmt.Println("El valor de y es:", *y)

	*x = 6
	fmt.Println("El valor modificado de x es:", *x)
	fmt.Println("El valor modificado de y es:", *y)
}
