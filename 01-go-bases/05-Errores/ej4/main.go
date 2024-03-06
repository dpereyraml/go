/*
Ejercicio 4 - Impuestos de salario #4

Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”,
para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
(el mensaje mostrado por consola deberá decir:

	“Error: the minimum taxable amount is 150,000 and the salary entered is: [salary]”,
	siendo [salary] el valor de tipo int pasado por parámetro)

		Ejercicio 3 - Impuestos de salario #3
		Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que,
		en reemplazo de “Error()”,  se implemente “errors.New()”.

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

var (
	// ErrSalary is less to limit
	ErrSalary = "Error: salary is less than 10000"
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
	salary = 210 // error
	// salary = 210000 // exito

	_, err := calculator(salary)

	if err != nil {
		if errors.Is(err, &CustomError{}) {
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
		err := fmt.Errorf(ErrSalary)
		return false, err
	}
	result = true
	return
}
