package main

import "fmt"

func main() {

	// leia um número inteiro e mostre o seu dobro, triplo e quadruplo
	var num1 int

	fmt.Print("digite um numero inteiro : ")
	fmt.Scan(&num1)

	// potencial número
	var calculodobro = num1 * 2
	var calculotriplo = num1 * 3
	var calculoquadruplo = num1 * 4
	println("o dobro do numero é: ", calculodobro, "o triplo do numero é: ", calculotriplo, "o quadruplo do numero é: ", calculoquadruplo)
}
