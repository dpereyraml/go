/*
	Ejercicio 3 - A qué mes corresponde
	Realizar una aplicación que reciba  una variable con el número del mes.
	Según el número, imprimir el mes que corresponda en texto.
	¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
	Ej: 7, Julio.
	Nota: Validar que el número del mes, sea correcto.
*/

package main

import "fmt"

var meses = map[int]string{ // es var porque Map se podria modificar y no deja que sera const
	1:  "Enero",
	2:  "Febrero",
	3:  "Marzo",
	4:  "Abril",
	5:  "Mayo",
	6:  "Junio",
	7:  "Julio",
	8:  "Agosto",
	9:  "Septiembre",
	10: "Octubre",
	11: "Noviembre",
	12: "Diciembre",
}

func main() {

	var indexMes int
	fmt.Println("Por favor ingrese un numero del 1 al 12")
	fmt.Scanln(&indexMes)
	// ><

	for indexMes < 1 || indexMes > 12 {
		fmt.Println("Intente nuevamente. Por favor ingrese un numero del 1 al 12")
		fmt.Scanln(&indexMes)
	}
	nombreMes := NombrarMes(indexMes)
	fmt.Println("El numero ingresado corresponde al mes de:02-estructurascontrol/ej3", nombreMes)
}

func NombrarMes(index int) string {
	return meses[index]
}
