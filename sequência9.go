package main

import "fmt"

func main() {
	//desconto=
	var preçopro float64

	fmt.Print("digite o valor do produto: ")
	fmt.Scanln(&preçopro)

	//calculo do desconto
	calculo := preçopro - (preçopro/100)*10
	fmt.Printf("o preço do produto foi reajustado:%.2f ", calculo)
}
