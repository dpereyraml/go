/*
Ejercicio 2 - Impuestos de salario #2

En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: salary is less than 10000"
y lanzalo en caso de que “salary” sea menor o igual a  10000.

La validación debe ser hecha con la función “Is()” dentro del “main”.
*/
package main

import (
	"errors"
	"fmt"
)

const (
	limit = 10000
)

/*
var (

	// ErrSalary is less to limit
	ErrSalary = "Error: salary is less than 10000"

)
*/
var ErrSalary = &CustomError{"Error: salary is less than 10000"}

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message //
	// return fmt.Sprintf("Error: the salary entered does not reach the taxable minimum: %s")
}

func main() {
	var salary int
	salary = 210

	_, err := calculator(salary)

	if err != nil {
		if errors.Is(err, ErrSalary) {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Must pay tax")

}

func calculator(salary int) (result bool, err error) {
	if salary < limit {
		return false, ErrSalary
		// return false, &CustomError{"Error: salary is less than or equal to 10000"}
	}
	result = true
	return
}
