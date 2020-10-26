package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	var numero2 int = 10000000000000
	fmt.Println(numero)
	fmt.Println(numero2)

	// unit nunca terá sinal
	const numero3 uint = 1000

	const numeroReal1 float32 = 1.1
	const numeroReal2 float64 = 1.2

	// um caracter entre '' retorna o valor dele na tabela ascii
	char := 'A'
	fmt.Println(char)

	//tipo erro
	var erro error = errors.New("Tipo erro")
	fmt.Println(erro)

}
