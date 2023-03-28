package main

import "fmt"

func main() {

	//calcule o imc
	var peso float64
	var altura float64

	fmt.Print("digite o seu peso?:")
	fmt.Scan(&peso)
	fmt.Print("digite sua altura?:")
	fmt.Scan(&altura)

	//calculo do imc
	calculoimc := peso / (altura * altura)
	fmt.Printf("o seu imc é: %.2f", calculoimc)

}
