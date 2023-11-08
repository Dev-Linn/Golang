package main

import "fmt"

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

func calc(n1, n2 int8) (int8, int8) {
	adicao := n1 + n2
	sub := n1 - n2

	return adicao, sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var a = func(txt string) string {
		fmt.Println(txt)

		return txt

	}

	retorno := a("Texto da função 1")
	fmt.Println(retorno)

	// Se eu substituir as variaveis (resultadosSoma, resultadoSub) por _ ele ignora um e da certo.
	resultadosSoma, resultadoSub := calc(10, 16)
	fmt.Println(resultadosSoma, resultadoSub)

	//teo
}
