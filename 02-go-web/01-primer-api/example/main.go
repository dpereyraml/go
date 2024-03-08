package example

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Conexion() {

	// server
	rt := chi.NewRouter()

	// - enpoinds
	rt.Get("/primer-test", func(w http.ResponseWriter, r *http.Request) {
		// code
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hola mundo desde mi primer api"))
	})

	// run
	if err := http.ListenAndServe(":8080", rt); err != nil {
		fmt.Println(err)
	}
}
