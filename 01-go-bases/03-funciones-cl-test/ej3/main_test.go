/*
	Ejercicio 3 - Test del salario
	La empresa marinera no está de acuerdo con los resultados obtenidos en los cálculos de los salarios,
	por ello nos piden realizar una serie de tests sobre nuestro programa.
	Necesitaremos realizar las siguientes pruebas en nuestro código:

	Calcular el salario de la categoría “A”.
	Calcular el salario de la categoría “B”.
	Calcular el salario de la categoría “C”.
*/

package main

import (
	"testing"
)

func TestCalculaLowSalaryA(t *testing.T) {
	// arrange
	minutosTrabajados := 180
	categoria := CatA
	expectedSalary := 3000.0
	// act
	obtainValue := CalcularSalario(minutosTrabajados, categoria)

	// assert
	if obtainValue != expectedSalary {
		t.Errorf("Error: obtainValue %.2f != expectedSalary %.2f", obtainValue, expectedSalary)

	}

}
