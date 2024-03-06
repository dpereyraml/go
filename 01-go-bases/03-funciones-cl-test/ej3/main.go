/*
	Ejercicio 3 - Calcular salario

	Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

	Categoría C, su salario es de $1.000 por hora.
	Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
	Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.

	Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.

*/

package main

import "fmt"

const ( // se podria hacer un orchestador
	CatA         = "A"
	CatB         = "B"
	CatC         = "C"
	SalaryCatA   = 1000
	SalaryCatB   = 1500
	SalaryCatC   = 3000
	MinutesXHour = 60.0
)

func main() {
	var minutosTrabajados int = 160

	salario := CalcularSalario(minutosTrabajados, CatA)

	fmt.Println(salario)
}

/* func CalcularSalario(cantHoras float64, salaryCat float64) float64 {

	return cantHoras * salaryCat
} */

func CalcularSalario(minutosTrabajados int, categoria string) float64 {
	horasTrabajadas := float64(minutosTrabajados) / MinutesXHour

	var salarioBase float64
	switch categoria {
	case CatA:
		salarioBase = SalaryCatA
	case CatB:
		salarioBase = SalaryCatB
	case CatC:
		salarioBase = SalaryCatC
	default:
		fmt.Println("Categoría no válida")
		return 0
	}

	var salarioTotal float64
	switch categoria {
	case "A":
		salarioTotal = horasTrabajadas * salarioBase
	case "B":
		salarioTotal = horasTrabajadas*salarioBase + (0.2 * salarioBase)
	case "C":
		salarioTotal = horasTrabajadas*salarioBase + (0.5 * salarioBase)
	}

	return salarioTotal
}
