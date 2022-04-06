package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	fmt.Println(&x)
	y := &x
	fmt.Println(y)  // Imprime la direccion de memoria de x.
	fmt.Println(*y) // Imprime el valor que esta almacenado en la direccion de x.
}
