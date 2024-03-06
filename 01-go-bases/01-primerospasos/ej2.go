// Ejercicio 2 - Clima

// Una empresa de meteorología quiere una aplicación donde pueda tener la temperatura,
// humedad y presión atmosférica de distintos lugares.

// Declará 3 variables especificando el tipo de dato, como valor deben tener la temperatura,
// humedad y presión de donde te encuentres.
// Imprimí los valores de las variables en consola.
// ¿Qué tipo de dato le asignarías a las variables?
package main

import "fmt"

func main() {
	var temperatura float64
	var humedad float64
	var presion float64

	temperatura = 29.3
	humedad = 44.1
	presion = 1013.25

	fmt.Println("Temperatura:", temperatura, "°C")
	fmt.Println("Humedad:", humedad, "%")
	fmt.Println("Presión Atmosférica:", presion, "mb")
}
