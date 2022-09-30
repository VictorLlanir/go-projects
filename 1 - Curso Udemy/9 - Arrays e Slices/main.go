package main

import "fmt"

func main() {
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"Pos1", "Pos2", "Pos3", "Pos4", "Pos5"}
	fmt.Println(array2)

	slice := []int{10, 11, 12, 13, 14}
	fmt.Println(slice)

	slice = append(slice, 15)
	fmt.Println(slice)

	slice2 := array2[1:3] // O primeiro parâmetro (1) é inclusivo; o segundo (3) é exclusivo
	fmt.Println(slice2)
}
