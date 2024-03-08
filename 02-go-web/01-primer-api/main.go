package main

import (
	"fmt"
	"primer-api/example"
	"primer-api/native"
	"primer-api/post"
)

func main() {

	var selection uint8

	fmt.Println("1: example - 2: nativo - 3: post greetings chi - 4: post greetings native")
	fmt.Scanln(&selection)

	switch selection {
	case 1:
		fmt.Println("endpoints: localhost:8080/primer-test")
		example.Conexion() // Call the "example.main()" function
	case 2:
		fmt.Println("endpoints: localhost:8080/ping")
		native.Native()
	case 3:
		fmt.Println("Chi endpoints: localhost:8080/greetings")
		post.PostChi()
	case 4:
		fmt.Println("Native endpoints: localhost:8080/greetings")

	default:
		fmt.Println("endpoints: localhost:8080/primer-test")
	}
}
