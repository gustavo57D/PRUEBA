package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Bienvenido a mi servidor!")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    fmt.Println("Servidor iniciado en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func main() {
    handleRequests()
}
