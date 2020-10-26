package main

import "fmt"

func main() {
	// maneiras declarar e atribuir valores a variavel
	var variavel1 string = "maneira explícita"
	variavel2 := "maneira implicita"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "aaaaa"
		variavel4 string = "bbbbb"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "ccccc", "ddddd"

	fmt.Println(variavel5, variavel6)

	const constante1 string = "const1"

	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
