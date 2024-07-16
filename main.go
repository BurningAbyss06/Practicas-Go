package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Alumno struct {
	Nombre string
	Notas  []float64
}

func LeerCvs(nombreArchivo string) ([]Alumno, error) {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err
	}
	lector := csv.NewReader(archivo)
	lector.Comma = ';'

	registros, err := lector.ReadAll()
	if err != nil {
		return nil, err
	}
	var alumnos []Alumno

	for _, registo := range registros {
		nombre := registo[0]
		var notas []float64
		for _, notaStr := range registo[1:] {
			nota, err := strconv.ParseFloat(notaStr, 64)
			if err != nil {
				return nil, err
			}
			notas = append(notas, nota)
		}
		alumnos = append(alumnos, Alumno{Nombre: nombre, Notas: notas})
	}
	return alumnos, nil
}

func CalcularPromedio(notas []float64) float64 {
	var suma float64
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

func CalcularMediaAritmetica(alumnos []Alumno) float64 {
	var suma float64
	var cantidadNotas int
	for _, alumno := range alumnos {
		for _, nota := range alumno.Notas {
			suma += nota
			cantidadNotas++
		}
	}
	return suma / float64(cantidadNotas)
}

func EscribirCvs(nombreArchivo string, alumnos []Alumno) error {
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()
	escritor := csv.NewWriter(archivo)
	defer escritor.Flush()
	for _, alumno := range alumnos {
		promedio := CalcularPromedio(alumno.Notas)
		registro := []string{alumno.Nombre, fmt.Sprintf("%.2f", promedio)}
		if err := escritor.Write(registro); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	alumnos, err := LeerCvs("notas.csv")
	if err != nil {
		fmt.Println("Error al leer el Csv:", err)
		return
	}
	mediaAritmetica := CalcularMediaAritmetica(alumnos)
	fmt.Printf("La media aritmetica de todos los alumnos es: %2f\n", mediaAritmetica)
	err = EscribirCvs("promedios.csv", alumnos)
	if err != nil {
		fmt.Println("Error al escribir en el cvs:", err)
		return
	}
	fmt.Println("Se pudo crear el cvs con exito")
}
