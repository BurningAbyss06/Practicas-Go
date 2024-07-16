package main

import "fmt"

func mapas() {
	Map1 := map[string]string{
		"Wilman": "Grandett",
		"Angyee": "Marin",
		"Rafael": "Blanco",
		"Pedro":  "Kys",
	}

	fmt.Println("Mapa inicial: ", Map1)

	fmt.Println("El apellido de mi mejor amiga es: ", Map1["Angyee"])

	Map1["Marcello"] = "Servitad"

	fmt.Println("Mapa con cambios de agregado: ", Map1)

	delete(Map1, "Pedro")
	fmt.Println("Mapa con cambios de eliminado: ", Map1)

}
