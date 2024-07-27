package main

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Titulo       string `json:"titulo"`
	NumeroVideos int    `json:"numero_videos"`
}

type Cursos []Curso

func Json() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cursos := Cursos{
			Curso{"Curso de Go", 30},
			Curso{"Curso de Python", 20},
			Curso{"Curso de C#", 23},
		}
		json.NewEncoder(w).Encode(cursos)
		//curso1 := Curso{"Curso de Go", 30}
		//json.NewEncoder(w).Encode(curso1)
	})
	http.ListenAndServe(":8000", nil)
}
