package main

import "fmt"

func main() {

	//soma de 3 numeros inteiros
	var num1 int
	var num2 int
	var num3 int

	fmt.Print("digite um numero inteiro:")
	fmt.Scan(&num1)
	fmt.Print("digite outro numero inteiro:")
	fmt.Scan(&num2)
	fmt.Print("digite outro numero inteiro:")
	fmt.Scan(&num3)

	//calculo da soma
	calculo := num1 + num2 + num3
	println("o resultado da soma é:", calculo)

}
