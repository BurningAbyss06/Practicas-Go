package main

import (
	"fmt"
	"os"
)

func LeerTodoElArhivo() {
	file_data, err := os.ReadFile("./hola.txt")

	if err != nil {
		fmt.Println("Ocurrio un error", err)
	}
	fmt.Println(string(file_data))

}
