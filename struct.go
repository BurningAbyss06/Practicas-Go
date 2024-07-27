package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (usuario User) Nombre_Completo() string {
	return usuario.nombre + " " + usuario.apellido
}

/*
	func (this User) cambio_nombre(n string) {
		this.nombre = n
	}
*/
func (this *User) cambio_nombre1(n string) {
	this.nombre = n
}

type Persona struct {
	User
}

func estructuras() {
	/*persona1 := new(User)
	persona1.nombre = "Wilman"
	persona1.apellido = "Grandett"

	fmt.Println(persona1.Nombre_Completo())

	persona1.cambio_nombre("Pablo")
	fmt.Println("Cambio nombre 1", persona1.Nombre_Completo())
	persona1.cambio_nombre1("Pablo")
	fmt.Println("Cambio nombre 2", persona1.Nombre_Completo())*/

	n := "Wilmnan"
	a := "Grandett"
	persona1 := new(Persona)
	persona1.nombre = n
	persona1.apellido = a

	persona2 := User{nombre: "Angyee", apellido: "Marin"}

	fmt.Println(persona1.Nombre_Completo())
	fmt.Println(persona2.Nombre_Completo())

}
