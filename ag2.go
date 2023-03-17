package main

import "fmt"

func main() {

	var x1 int

	fmt.Print("qual o número?")
	fmt.Scan(&x1)
	dobro := (x1 * 2)
	triplo := (x1 * 3)
	quadruplo := (x1 * 4)
	fmt.Println(dobro, triplo, quadruplo)
}
