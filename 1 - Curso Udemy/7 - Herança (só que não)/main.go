package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type Estudante struct {
	Pessoa
	curso     string
	faculdade string
}

func main() {
	est := Estudante{Pessoa{"Victor", "Llanir", 24}, "Engenharia", "FIAP"}
	fmt.Println(est.nome)
}
