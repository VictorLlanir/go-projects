package main

import "fmt"

func main() {
	var texto string = "Meu texto"
	nome := "Meu nome é engraçado"
	fmt.Println(texto)
	fmt.Println(nome)

	var (
		textao    string = "teas"
		textinho  string = "da"
		numerinho int    = 1
	)

	fmt.Println(textao, textinho, numerinho)

	numerao, runa := 1000, 33
	fmt.Println(numerao, runa)

	const minhaUltraConstante string = "CONST"
	fmt.Println(minhaUltraConstante)
}
