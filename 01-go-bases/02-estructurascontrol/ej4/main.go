/*
	Ejercicio 4 - Qué edad tiene...
	Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
	Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	Por otro lado también es necesario:
	- Saber cuántos de sus empleados son mayores de 21 años.
	- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	- Eliminar a Pedro del mapa.
*/

package main

import (
	"fmt"
)

const limiteEdad = 21

var Employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	MayoresDeEdad()
	AgregarEmpleado("Federico", 25)
	QuitarEmpleado("Federico")
}

func MayoresDeEdad() {
	fmt.Println("Funcion MayoresDeEdad a Employees se esta ejecutando...")
	var count int

	for _, e := range Employees {
		if e > limiteEdad {
			count += 1
		}
	}
	fmt.Println("Cantidad de empleados mayores de", limiteEdad, "años es", count)
	fmt.Println("--------")
}

// ><
func AgregarEmpleado(nombreE string, edad int) {
	fmt.Println("Agrear elemento a Employees se esta ejecutando...")
	if edad < 1 {
		fmt.Println("No se puede agregar el empleado porque todavia no ha nacido")
	} else {
		Employees[nombreE] = edad
		//fmt.Println(Employees)
		ImprimirEmployee()
	}
	return
}

func QuitarEmpleado(nombre string) {
	fmt.Println("Quitar elemento a Employees se esta ejecutando...")
	delete(Employees, "Brenda")
	//fmt.Println(Employees)
	ImprimirEmployee()
	return
}

func ImprimirEmployee() {
	for nombre, edad := range Employees {
		fmt.Println(nombre, edad)
	}
	fmt.Println("--------")
}
