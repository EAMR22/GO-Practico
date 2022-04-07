package main

import "fmt"

func main() {
	m1 := make(map[string]int) // map es una estructura de llave valor.

	m1["a"] = 8          // EL string es la llave y el int que es el valor que apunta la llave.
	fmt.Println(m1)      // map[a:8].
	fmt.Println(m1["a"]) // Imprime el valor de 8.
}
