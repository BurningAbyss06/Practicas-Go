package main

import (
	"fmt"
	"strings"
)

func bucles() {
	/*var sum int = 0
	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			sum = sum + i
		}
	}
	fmt.Println("La suma de los primero 100 numeros impares es:", sum)*/
	/*Map1 := map[string]string{
		"Wilman": "Grandett",
		"Angyee": "Marin",
		"Rafael": "Blanco",
		"Pedro":  "Kys",
	}

	for k, v := range Map1 {
		fmt.Println("Nombre: " + k + " Apellido: " + v)
	}*/

	var fruta string = "pera"

	for {
		if fruta == "naranja" {
			fmt.Println("La fruta que eligio es: ", fruta)
			break
		}
		fmt.Println("Introducir una fruta")
		fmt.Scan(&fruta)
		fruta = strings.ToLower(fruta)
	}
}
