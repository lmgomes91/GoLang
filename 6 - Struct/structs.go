package main

import "fmt"

type endereco struct {
	rua    string
	numero int16
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("structs")

	var u usuario
	u.nome = "Lucas"
	u.idade = 28
	fmt.Println(u)

	endereco1 := endereco{rua: "Campestre", numero: 184}

	usuario2 := usuario{"Lucas", 27, endereco1}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Lucas"}
	fmt.Println(usuario3)

}
