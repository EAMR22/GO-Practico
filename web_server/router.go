package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServerHTTP(w http.ResponseWriter, request *http.Request) { // parametros: el primero es el escritor, el segundo es el request en donde viene la informacion
	fmt.Fprintf(w, "Hello World!") // Fprintf utiliza un escritor, que recive w que es el escritor asignado, y el mensaje que queremos mostrar.
}
