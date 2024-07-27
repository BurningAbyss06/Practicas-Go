package main

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(name string) {
	letras := strings.Split(name, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}

func rutinas() {
	go MiNombreLento("Wilman")
	var wait string
	fmt.Scan(&wait)

}
