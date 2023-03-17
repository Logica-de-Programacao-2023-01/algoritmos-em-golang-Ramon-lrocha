package main

import "fmt"

func main() {
	var x1, x2, x3 float64
	fmt.Print("qual o valor do primeiro numero?")
	fmt.Scan(&x1)
	fmt.Print("qual o valor do segundo numero?")
	fmt.Scan(&x2)
	fmt.Print("qual o valor do terceiro numero?")
	fmt.Scan(&x3)
	resultado := (x1 + x2 + x3)
	fmt.Println("a media é", resultado)

}
