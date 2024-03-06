/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.

Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará además un 10 % (27% en total).

*/

package main

import "fmt"

const (
	porcentajeMenor = 0.17
	porcentajeMayor = 0.27
)

func main() {
	var sueldoBr float64

	fmt.Println("Ingrese sueldo Bruto")
	fmt.Scanln(&sueldoBr)

	impuesto := CalculaImpuesto(sueldoBr)

	var sueldoNeto = sueldoBr - impuesto

	fmt.Println("Su sueldo Neto es ", sueldoNeto, "Tiene impuestos por un valor de", impuesto)
}

func CalculaImpuesto(sueldo float64) float64 {
	if sueldo <= 50000 {
		return 0
	} else if sueldo > 150000 {
		return sueldo * porcentajeMayor
	} else {
		return sueldo * porcentajeMenor
	}
}
