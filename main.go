package main

import "fmt"

func main(){
	const USDtoEUR = 0.87
	const USDtoRUB = 80.05
	var EUR float64 = 80
	EURtoRUB := EUR / USDtoEUR * USDtoRUB

	fmt.Printf("EUR in RUB: %.2f", EURtoRUB)
}