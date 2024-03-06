/*
Ejercicio 3 - Declaración de variables


Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.

Necesita ayuda para:

Detectar cuáles de estas variables que declaró el alumno son correctas.
Corregir las incorrectas.
   var 1firstName string

   var lastName string

   var int age

   1lastName := 6

   var driver_license = true

   var person height int

   childsNumber := 2

*/

package main

func main() {

	var lastName string

	var int age

	lastName := 6 // tenia error porque comenzaba con 1 y no estaba declarada

	var driver_license = true

	var person_height int // entiendo que era una variable con nombre compuesto y la separe con guion bajo

	childsNumber := 2

}
