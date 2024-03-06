package main

import (
	"testing"
)

func TestCalculaPromedio(t *testing.T) {
	// arrange
	arregloN := []float64{1.1, 2.2, 3.3, 4.4, 5.5} // slice
	expectedProm := 3.3
	// act
	obtainValue := CalcaularPromedio(arregloN)

	// assert
	if obtainValue != expectedProm {
		t.Errorf("Error")
	}

}
