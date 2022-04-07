package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
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

	lista.agregarALista(t3)

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index", i, "Nombre", lista.tasks[i].nombre)
	}

	for index, tarea := range lista.tasks { // range devuelve 2 valores y a la vez crea un iterador que recorre el ciclo.
		fmt.Println("Index", index, "Nombre", tarea.nombre)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break // break lo que hace es parar el ciclo en esa condicion.
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // Simplemente rompe la iteracion, pero el ciclo continua.
		}
		fmt.Println(i)
	}
}
