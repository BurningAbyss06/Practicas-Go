package main

import "fmt"

func condicionales() {
	/*var edad int = 15

	if edad >= 18 {
		fmt.Println("Eres mayor de eddad")
	} else {
		fmt.Println("No eres mayor de edad, intentalo en unos años")
	}*/
	var deporte string
	fmt.Println("Ingrese el deporte que esta jugando Pepito ")
	fmt.Scan(&deporte)
	switch deporte {
	case "Futbol":
		fmt.Println("Pepito esta jugando al", deporte)
	case "Volibol":
		fmt.Println("Pepito esta jugando al", deporte)
	case "Basket":
		fmt.Println("Pepito esta jugando al", deporte)
	default:
		fmt.Println("Pepito no esta jugando a nada")
	}
}
