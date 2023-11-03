package main

import (
	"fmt"
	"moodulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello word!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("lincoln@gmail.com")
	fmt.Println(erro)

	if erro == nil {
		fmt.Println("O email é válido")
	} else {
		fmt.Println("O email é inválido")
	}

}
