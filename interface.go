package main

import "fmt"

type animal interface {
	mover() string
}

type perro struct{}

type pez struct{}

type ave struct{}

func (perro) mover() string { // La clase perro se convierte en una interfaz animal.
	return "Soy un perro y estoy caminando"
}

func (pez) mover() string {
	return "Soy un pez y estoy nadando"
}

func (ave) mover() string {
	return "Soy un ave y estoy volando"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	p := perro{}
	moverAnimal(p)

	pe := pez{}
	moverAnimal(pe)

	pa := ave{}
	moverAnimal(pa)
}
