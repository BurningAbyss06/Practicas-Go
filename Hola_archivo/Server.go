package main

import (
	"net/http"
)

func htmlWeb() {
	//Una forma de hacerlo pero es mas insegura
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, r.URL.Path[1:])

	// })
	// http.ListenAndServe(":8000", nil)

	file_server := http.FileServer(http.Dir("Public"))
	http.Handle("/", http.StripPrefix("/", file_server))
	http.ListenAndServe(":8000", nil)
}
