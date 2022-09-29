package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func calculosMatematicos(n1, n2 float32) (float32, float32, float32, float32) {
	return n1 + n2, n1 - n2, n1 * n2, n1 / n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Texto de exemplo")
	fmt.Println(result)

	somaComplexa, _, multiplicacao, divisao := calculosMatematicos(10, 15) // O caracter _ serve para descartar um valor que não utilizaremos
	fmt.Println(somaComplexa, multiplicacao, divisao)
}
