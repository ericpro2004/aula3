package main

import "fmt"

func main() {
	//cálculo báscio da área do retângulo.

	var base int
	var lado int

	fmt.Print("qual é a base do retângulo")
	fmt.Scanln(&base)
	fmt.Print("qual é a lado do retângulo")
	fmt.Scanln(&lado)

	var area_rt int = base * lado

	fmt.Println("O valor da base do retângulo é: ", base, "O valor do lado do retângulo é: ", lado, "O resultado da area_rt é: ", area_rt)

}
