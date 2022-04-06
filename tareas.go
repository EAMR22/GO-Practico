package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) { // Lo enviamos con * para evitar hacer modificaciones a valores copia.
	t.tasks = append(t.tasks, tl) // El 1 parametro toma el slice y el 2 parametro toma el valor que se agregara al slice.
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...) // Los ... es porque el 2 parametro no es un slice.
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:      "Completar mi curso de go",
		descripcion: "Completar mi curso de go de Platzi esta semana",
	}

	t2 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python de Platzi esta semana",
	}

	t3 := &task{
		nombre:      "Completar mi curso de nodeJS",
		descripcion: "Completar mi curso de nodeJS de Platzi esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	fmt.Println(lista.tasks[0])
	lista.agregarALista(t3)
	fmt.Println(len(lista.tasks)) // Imprime 3
	lista.eliminarDeLista(1)
	fmt.Println(len(lista.tasks)) // Imprime 2
}
