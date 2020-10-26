package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("aprendendo a trabalhar com modulos")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("lgomes@post.com")
	fmt.Println(erro)
}
