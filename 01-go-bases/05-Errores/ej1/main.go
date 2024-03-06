/*
	Ejercicio 1 - Impuestos de salario #1
	En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
	Creá un error personalizado con un struct que implemente “Error()” con el mensaje
	“Error: the salary entered does not reach the taxable minimum" y lanzalo en caso de que “salary” sea menor a 150.000.
	De lo contrario, tendrás que imprimir por consola el mensaje “Must pay tax”.

*/

package main

import "fmt"

const (
	limit = 150000
)

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message //
	// return fmt.Sprintf("Error: the salary entered does not reach the taxable minimum: %s")
}

func main() {
	var salary int
	salary = 121000

	if salary < limit {
		err := &CustomError{"Error the salary entered does not reach the taxable minimum"}
		// err := &CustomError{}
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Must pay tax")

}
