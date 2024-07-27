package main

import "fmt"

func canal() {
	channel := make(chan string)

	go func(channel chan string) {
		for {
			var name string
			fmt.Scanln(&name)
			channel <- name
		}
	}(channel)

	msg := <-channel

	fmt.Println("Esto es lo que recibi del canal:", msg)

}
