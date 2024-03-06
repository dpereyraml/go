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

	sdfile, err := os.Open("customer.txt")

	if err != nil {
		err = fmt.Errorf("The indicated file was not found or is damaged%w", err)
		panic(err)
	}
	srdfile, err := os.ReadFile(sdfile.Name())
	fmt.Println(string(srdfile))
}
