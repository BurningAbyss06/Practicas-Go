package main

import (
	"bufio"
	"fmt"
	"os"
)

func AbrirArchivo() {
	archivo, err := os.Open("./holas.txt")
	defer func() {
		archivo.Close()
		fmt.Println("Se efecucto el cierre del archivo")
		r := recover()
		fmt.Println(r)
	}()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan() {
		i++
		linea := scanner.Text()
		fmt.Println("Estoy en la linea:", i)
		fmt.Println(linea)
	}

}

func panico() {
	AbrirArchivo()
	fmt.Println("")
	fmt.Println("Todo salio bien")
}
