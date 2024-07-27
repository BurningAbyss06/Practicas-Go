package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hay una nueva peticion :D")
	io.WriteString(w, "hola mundo soy una pag web")
}

func webss() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
