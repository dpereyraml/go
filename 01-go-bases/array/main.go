package main

import "fmt"

func main() {
	// arreglo
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// slice

	var slice = []string{"hola", "mundo"}
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)

	// slice con make

	aa := make([]int, 5)
	fmt.Println(len(aa))

	// slice obtenido de un rango

	primes := []int{2, 5, 6, 7, 8, 23, 45}
	fmt.Println(primes[1:4])

	longitud := len(primes)
	capacidad := cap(primes)

	fmt.Println(longitud, capacidad)

	// agregar elementos a slice

	primes = append(primes, 88, 22, 44)

	fmt.Println(primes)

	// MAPS permiten crear variables tipo clave-valor
	// dos formas de definir
	myMapDos := make(map[string]int)
	myMapDos = map[string]int{"Benjamin": 12, "Mariana": 33}
	fmt.Println(myMapDos)
	fmt.Println("Benajmin")
	fmt.Println(myMapDos["Benajmin"])

	myMap := map[string]int{"Benjamin": 12, "Mariana": 33}
	myMap["David"] = 35      // agrego elemento
	delete(myMap, "Mariana") // borrar elemento
	fmt.Println(myMap)
	fmt.Println(myMap["Benajmin"])

	// recorrer elementos
	for key, element := range myMapDos {
		fmt.Println("Key: ", key, " => Element ", element)
	}
}
