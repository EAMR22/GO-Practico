package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Creamos un scanner para que el usuario pueda digitar.
	scanner.Scan()
	operacion := scanner.Text() // Convierte el scanner a string.
	fmt.Println(operacion)
	valores := strings.Split(operacion, "+") // Con Split separa la cadena de texto.
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])
	operador1, _ := strconv.Atoi(valores[0]) //Atoi devuelve 2 valores.
	operador2, _ := strconv.Atoi(valores[0]) // El "_" nos dice que no utilizaremos la otra variable.
	fmt.Println(operador1 + operador2)

}
