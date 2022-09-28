package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos de dados
	// Inteiros: int, int8, int16, int32 (rune), int64 - o número deles indica a quantidade de bits dentro desse inteiro. O int sozinho usa a arquitetura do computador como base
	// Inteiros sem sinal: uint, uint8 (byte), uint16, uint32, uint64 - as convenções são as mesmas de cima
	// Reais: float (não pode ser declarado explicitamente) float32, float64
	// String (cadeia de caracters): string. Em Go sempre se usa aspas duplas.
	// Char: Go não tem char. Char se tornará um número da tabela ASCII, referente àquele caracter

	// Valor zero: é o valor inicial de uma variável, caso ela não seja inicializada

	// Booleano: true ou false
	// Error: Seu valor zero é <nil>

	var erro error = errors.New("erro interno")
	fmt.Println(erro)
}
