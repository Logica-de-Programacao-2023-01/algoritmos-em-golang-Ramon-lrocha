package main

import "fmt"

func main() {

	var peso float64
	var altura float64

	fmt.Print("qual o seu peso?")
	fmt.Scan(&peso)
	fmt.Print("qual a sua altura?")
	fmt.Scan(&altura)
	resultado := peso / altura * altura
	fmt.Println("seu imc é ", resultado)

}
