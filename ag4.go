package main

import "fmt"

func main() {
	var x, y, z float64

	fmt.Print("qual os número?")
	fmt.Scan(&x, &y, &z)

	resultado := ((x * 2) + (y * 3) + (z * 5)) / 10

	fmt.Println("essa é a média", resultado)

}
