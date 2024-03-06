/*
Ejercicio 2 - Préstamo


Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.

Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

package main

import "fmt"

func main() {
	var edad uint8
	var trabajo bool
	var antiguedad float64
	var sueldo float64

	for {
		fmt.Println("Ingrese su edad en años")
		fmt.Scanln(&edad)
		if edad > 0 {
			break
		}
	}

	for {
		fmt.Println("Ingrese 1 si trabaja, cualquier otro valor sera tomado como falso")
		fmt.Scanln(&trabajo)
		if trabajo == false || trabajo == true {
			break
		}
	}

	for {
		fmt.Println("Ingrese su antiguedad laboral, puede expresarla en años y meses, ejemplo 1 año y medio es 1.5")
		fmt.Scanln(&antiguedad)
		if antiguedad >= 0 {
			break
		}
	}
	for {
		fmt.Println("Ingrese su sueldo")
		fmt.Scanln(&sueldo)
		if sueldo >= 0 {
			break
		}
	}

	fmt.Println(edad)
	fmt.Println(trabajo)
	fmt.Println(antiguedad)

	if edad > 22 && trabajo && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Felicidades, cumples con todos los requisitos y puedes acceder a un préstamo sin intereses.")
		} else {
			fmt.Println("Cumples con los requisitos mínimos para acceder a un préstamo, pero tendrás que pagar intereses.")
		}
	} else {
		fmt.Println("Lo siento, no cumples con los requisitos necesarios para acceder a un préstamo.")
	}
}
