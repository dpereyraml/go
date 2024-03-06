/*
	Ejercicio 2 - Calcular promedio
	Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
	Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio.
	No se pueden introducir notas negativas.
*/

package main

import "fmt"

func main() {
	arregloN := []float64{1.1, 2.2, 3.3, 4.4, 5.5} // slice

	promedio := CalcaularPromedio(arregloN)

	fmt.Println(promedio)
}

func CalcaularPromedio(arregloN []float64) float64 {
	sumatoria := 0.0
	cantidad := float64(len(arregloN))
	for _, v := range arregloN {
		sumatoria += v
	}
	return sumatoria / cantidad
}
