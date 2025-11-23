package main

import "fmt"

func userInput() (string, string, float64) {
	var valueFrom string
	var valueTo string
	var value float64

	fmt.Print("Из какой валюты:")
	fmt.Scan(&valueFrom)

	fmt.Print("В какую валюту:")
	fmt.Scan(&valueTo)
	
	fmt.Print("Сколько:")
	fmt.Scan(&value)

	return valueFrom, valueTo, value
}

func calculateValue(valueFrom string, valueTo string, value float64) {

}

func main(){
	const USDtoEUR = 0.87
	const USDtoRUB = 80.05
	var EUR float64 = 80
	EURtoRUB := EUR / USDtoEUR * USDtoRUB

	fmt.Printf("EUR in RUB: %.2f", EURtoRUB)

	valueFrom, valueTo, value := userInput();
	calculateValue(valueFrom, valueTo, value)
}