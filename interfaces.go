package main

import "fmt"

type ProcesosUser interface {
	Permisos() int
	Nombre() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}
func (this Admin) Nombre() string {
	return this.nombre
}

func Autentificacion(user ProcesosUser) string {
	if user.Permisos() > 4 {
		return user.Nombre() + " tiene permisos de Administrador"
	}
	return ""
}

func interfaz() {

	//var admin = new(Admin)
	usuarios := []ProcesosUser{Admin{"Wilman"}, Admin{"Angyee"}}

	for _, usuario := range usuarios {
		fmt.Println(Autentificacion(usuario))
	}

}
