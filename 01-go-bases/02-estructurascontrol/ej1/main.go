/*
Ejercicio 1 - Letras de una palabra
La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:

Crear una aplicación que reciba una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
Luego, imprimir cada una de las letras.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var palabra string
	palabra = "Palabra"
	// fmt.Println((palabra[0])) // imprime el ascii
	// fmt.Println(len(palabra)) // imprime longitut
	// fmt.Println(string(palabra[0]))  // imprime P - la primer letra de la palabra

	fmt.Println("Ingrese un valor: ")
	fmt.Scanln(&palabra)
	// fmt.Println(palabra)
	var counter int

	fmt.Println("A deletrear estas ", len(palabra), " Letras")
	for counter < len(palabra) {
		time.Sleep(time.Second * 1)
		fmt.Println(string(palabra[counter]))
		counter++
	}
}
