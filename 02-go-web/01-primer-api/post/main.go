/*
Ejercicio 2 - Manipulando el body

Vamos a crear un endpoint llamado /greetings.
Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en texto “Hello + nombre + apellido”

El endpoint deberá ser de método POST
Se deberá usar el package JSON para resolver el ejercicio
La respuesta deberá seguir esta estructura: “Hello Andrea Rivas”
La estructura deberá ser como esta:

	{
		“firstName”: “Andrea”,
		“lastName”: “Rivas”
	}
*/

package post

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func PostNative() {
	//
}

func PostChi() {

	// server
	rt := chi.NewRouter()

	// - enpoinds
	rt.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		// code
		body, err := ioutil.ReadAll(r.Body) // usar os.ReadAll(r.Body) para leer el body ???
		if err != nil {
			http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
			return
		}

		fmt.Println("body", string(body))
		pers, err := decodingJson(string(body)) // usar json.newDecoder(r.Body).Decode(&pers) para decodificar el body
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		fmt.Print(pers.FirstName)
		response := "Hello " + pers.FirstName + " " + pers.LastName
		w.Write([]byte(response))
		/* w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hola mundo desde mi primer api")) */
	})

	// run
	if err := http.ListenAndServe(":8080", rt); err != nil {
		fmt.Println(err)
	}
}

func decodingJson(body string) (person Person, err error) {
	// Deserializar el JSON en una estructura Person
	fmt.Println("body in decodingJson", body)

	err = json.Unmarshal([]byte(body), &person)
	//err = json.Unmarshal(body, &person)
	if err != nil {
		fmt.Println("Error al deserializar JSON:", err)
		return
	}

	return
}
