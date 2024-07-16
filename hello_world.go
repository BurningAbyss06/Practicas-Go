package main

import "fmt"

func hello() {

	var mystring string = "Hola Como estas"

	fmt.Println(mystring)

	var numero int = 60
	fmt.Println(numero)
	numero2 := 40
	fmt.Println(numero2)

	var listaJuegos = []string{"LoL", "Terraria", "Guardian Tales", "Eversoul"}
	fmt.Println("Soy lista 1", listaJuegos)

	listaJuegos2 := listaJuegos[:2]
	fmt.Println("Soy lista 2", listaJuegos2)

}
