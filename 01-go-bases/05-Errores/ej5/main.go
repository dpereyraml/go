/*
Ejercicio 5 -  Impuestos de salario #5

Vamos a hacer que nuestro programa sea un poco más complejo y útil.

1. Desarrollá las funciones necesarias para permitir a la empresa calcular:
a) Salario mensual de un trabajador según la cantidad de horas trabajadas.
- La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
- Dicha función deberá retornar más de un valor (salario calculado y error).
- En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de impuesto.
- En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error. El mismo tendrá que indicar “Error: the worker cannot have worked less than 80 hours per month”.
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
