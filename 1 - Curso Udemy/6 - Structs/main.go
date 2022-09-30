package main

import "fmt"

type Usuario struct {
	nome     string
	idade    uint8
	endereco Endereco
}

type Endereco struct {
	logradouro string
	numero     uint16
}

func main() {
	fmt.Println("----- STRUCTS -----")

	u := Usuario{"Victor Llanir", 24, Endereco{"Rua Jaguaribe", 624}}

	fmt.Println(u)
}
