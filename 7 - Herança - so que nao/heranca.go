package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança, só que não")
	p := pessoa{nome: "Lucas", sobrenome: "Gomes", idade: 28, altura: 177}
	e := estudante{pessoa: p, curso: "Engenharia de computação", faculdade: "UFOP"}

	fmt.Println(e)
	fmt.Println(e.nome)

}
