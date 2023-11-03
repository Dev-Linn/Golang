package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int = 10
	fmt.Println(numero)

	//uint = int sem sinal (unsygned) - Não aceita números negativos

	var numero2 uint32 = 12
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 1234
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	///////////////////////////////////////////////////////

	var numeroReal1 float32 = 12.3
	fmt.Println(numeroReal1)

	///////////////////////////////////////////////////////

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
