package main

import "fmt"

func main() {
	var nome string
	var peso string
	var idade string
	fmt.Print("qual é seu nome")
	fmt.Scan(&nome)
	fmt.Print("Qual é seu peso")
	fmt.Scan(&peso)
	fmt.Print("qual é sua idade")
	fmt.Scan(&idade)
	println("O seu nome é: ", nome, "O seu peso é: ", peso, "A sua idade é: ", idade)

}
