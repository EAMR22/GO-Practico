package main

import (
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

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}

func (r *Router) ServerHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exist := r.FindHandler(request.URL.Path)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return // Es una manera de decirle a la funcion que pare la ejecucion
	}

	handler(w, request)
}
