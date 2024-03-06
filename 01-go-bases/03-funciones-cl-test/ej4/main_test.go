/*
	Los profesores de la universidad de Colombia, entraron al programa de análisis de datos  de Google, el cual premia a los mejores estadísticos de la región. Por ello los profesores nos solicitaron comprobar el correcto funcionamiento de nuestros cálculos estadísticos. Se solicita la siguiente tarea:

	Realizar test para calcular el mínimo de calificaciones.
	Realizar test para calcular el máximo de calificaciones.
	Realizar test para calcular el promedio de calificaciones.
*/

package main

import (
	"testing"
)

func TestCalculaLowSalaryA(t *testing.T) {
	// arrange
	hours := 200.0
	salary := CatA
	expectedSalary := 200000.0
	// act
	obtainValue := CalcularSalario(hours, salary)

	// assert
	if obtainValue != expectedSalary {
		t.Errorf("Error")
	}

}
