package main

import "fmt"

func main() {

	//Cálculo do volume de uma caixa.
	var Profundidade float32
	var Base float32
	var Altura float32

	fmt.Print("quanto é a Profundidade da caixa?")
	fmt.Scan(&Profundidade)
	fmt.Print("quanto é a Base da caixa")
	fmt.Scan(&Base)
	fmt.Print("quanto é a Altura da caixa")
	fmt.Scan(&Altura)

	var volume_cx float32 = Altura * Base * Profundidade

	fmt.Println("A profundidade é: ", Profundidade, "A base é: ", Base, "A altura é", Altura, volume_cx)

}
