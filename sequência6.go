package main

import "fmt"

func main() {

	var salário float64

	fmt.Print("digite seu salário:")
	fmt.Scanln(&salário)

	//aumento
	calculo := (salário/100)*15 + salário
	fmt.Printf("seu salário foi reajustado para:%.2f ", calculo)
}
