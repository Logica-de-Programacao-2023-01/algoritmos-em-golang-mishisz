package main

import "fmt"

func main() {
	//peso

	var peso2 float64
	var peso3 float64
	var peso5 float64

	fmt.Print("digite o valor de peso2:")
	fmt.Scanln(&peso2)
	fmt.Print("digite o valor de peso3:")
	fmt.Scanln(&peso3)
	fmt.Print("digite o valor de peso5:")
	fmt.Scanln(&peso5)

	//calculo
	calculomedia := (peso2*2 + peso3*3 + peso5*5) / (2 + 3 + 5)
	fmt.Printf("o resultado é:%.2f ", calculomedia)
}
