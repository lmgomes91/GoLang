package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var valor int
	var ponteiro *int

	valor = 100
	ponteiro = &valor

	fmt.Println(valor, ponteiro, *ponteiro)
}
