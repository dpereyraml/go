/*
Ejercicio 1 - Prueba de Ping

Vamos a crear una aplicación web con el package net/http nativo de go, que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”

1.   El endpoint deberá ser de método GET
2.   La respuesta de “pong” deberá ser enviada como texto, NO como JSON
*/
package native

import (
	"net/http"
)

func Native() {
	// handler
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pong!"))
	})

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		println("Error starting server: ", err.Error())
		return
	}
}
