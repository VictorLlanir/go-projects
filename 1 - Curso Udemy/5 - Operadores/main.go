package main

import "fmt"

func main() {
	// Operadores Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	// Operadores de atribuição
	var texto string = "texto"
	nome := "meu nome"
	fmt.Println(texto, nome)

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores lógicos
	fmt.Println(true && true)
	fmt.Println(false || true)
	fmt.Println(!true)

	// Operadores unários
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 33
	numero *= -1
	numero /= 2
	fmt.Println(numero)
}
