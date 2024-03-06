/*
	Ejercicio 1 - Testear el impuesto del salario
	La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados al momento de depositar el sueldo de los mismos ahora nos solicitó validar que los cálculos de estos impuestos están correctos. Para esto nos encargaron el trabajo de realizar los test correspondientes para:

	Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
	Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
	Calcular el impuesto en caso de que el empleado gane por encima de $150.000.
*/

package main

import "testing"

func TestCalculaImpuestoLowerSalary(t *testing.T) {
	// arrange
	var salary float64 = 30000
	var expectedTax float64 = 0 // el impuesto esperado

	// act
	obtainTax := CalculaImpuesto(salary)

	// assert
	if obtainTax != expectedTax {
		t.Errorf("Error")
	}
}
