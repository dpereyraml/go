package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Ejecucion finalizada. Error:%v\n", r)
		}
	}()

	f, err := os.Open("customers.txt")
	defer f.Close()
	if err != nil {
		err = fmt.Errorf("The indicated file was not found or is damaged%w", err)
		panic(err)
	}
	srdfile, err := os.ReadFile(f.Name())
	if err != nil {
		err = fmt.Errorf("The indicated file was not found or is damaged%w", err)
		panic(err)
	}
	fmt.Println(string(srdfile))
	fmt.Println("Ejecucion finalizada")
}
