package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()
	canal := make(chan string)
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal) // go crea un nuevo hilo o un nuevo proceso, independientemente de main.
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal) // El canal esta pasando la data.
	}

	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoPaso)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "No esta disponible =(")
		canal <- servidor + " No se encuentra disponible" // Transmitimos la data al canal.
	} else {
		fmt.Println(servidor, "Esta funcionando normalmente =)")
		canal <- servidor + " Esta funcionando"
	}
}
