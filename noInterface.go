package main

import "fmt"

type perro struct{}

type pez struct{}

type ave struct{}

func (perro) camina() string {
	return "Soy un perro y estoy caminando"
}

func (pez) nada() string {
	return "Soy un pez y estoy nadando"
}

func (ave) vuela() string {
	return "Soy un ave y estoy volando"
}

func moverPerro(p perro) {
	fmt.Println(p.camina())
}

func moverPez(p pez) {
	fmt.Println(p.nada())
}

func moverAve(p ave) {
	fmt.Println(p.vuela())
}

func main() {
	p := perro{}
	moverPerro(p)

	pe := pez{}
	moverPez(pe)

	pa := ave{}
	moverAve(pa)
}
