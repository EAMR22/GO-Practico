package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() { // Utiliza el * para actualizar el nuevo valor que viene de la referencia.
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := &task{ // Ya no enviamos una copia sino una referencia del struct que creamos.
		nombre:      "Completar mi curso de go",
		descripcion: "Completar mi curso de go de Platzi esta semana",
	}
	fmt.Printf("%+v\n", t)

	t.marcarCompleta()
	t.actualizarDescripcion("Finalizar mi curso de go")
	t.actualizarNombre("Cmpletar mi curso cuanto antes")
	fmt.Printf("%+v\n", t)
}
