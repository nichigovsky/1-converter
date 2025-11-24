package main

import (
	"fmt"
)

const EUR = "EUR"
const USD = "USD"
const RUB = "RUB"

func isCorrectValue(valFrom, valTo string) (bool, string) {
	if valFrom == valTo {
		return false, "Валюты не могут совпадать"
	}

	isCorrect := valFrom == EUR || valFrom == USD || valFrom == RUB
	return isCorrect, "Введите допустимое значение: USD, RUB, EUR"
}

func isCorrectCurrency(val float64) (bool, string) {
	isCorrect := val > 0
	return isCorrect, "Введите число или значение больше нуля"
}

func userInput() (string, string, float64) {
	var valueFrom string
	var valueTo string
	var value float64

	fmt.Print("Из какой валюты: ")
	for {
		fmt.Scan(&valueFrom)
		isCorrect, str := isCorrectValue(valueFrom, valueTo)
		if !isCorrect {
			fmt.Println(str)
			continue
		}
		break
	}

	fmt.Print("В какую валюту: ")
	for {
		fmt.Scan(&valueTo)
		isCorrect, str := isCorrectValue(valueTo, valueFrom)
		if !isCorrect {
			fmt.Println(str)
			continue
		}
		break
	}
	
	fmt.Print("Сколько: ")
	for {
		fmt.Scan(&value)
		isCorrect, str := isCorrectCurrency(value)
		if !isCorrect {
			fmt.Println(str)
			continue
		}
		break
	}

	return valueFrom, valueTo, value
}

func calculateValue(valueFrom string, valueTo string, value float64) {
	const USDtoEUR = 0.87
	const USDtoRUB = 80.05

	switch {
	case valueFrom == USD && valueTo == RUB:
		fmt.Printf("Перевод из %s в %s - %.2f", valueFrom, valueTo, value * USDtoRUB)
	case valueFrom == RUB && valueTo == USD:
		fmt.Printf("Перевод из %s в %s - %.2f", valueFrom, valueTo, value / USDtoRUB)
	case valueFrom == USD && valueTo == EUR:
		fmt.Printf("Перевод из %s в %s - %.2f", valueFrom, valueTo, value * USDtoEUR)
	case valueFrom == EUR && valueTo == USD:
		fmt.Printf("Перевод из %s в %s - %.2f", valueFrom, valueTo, value / USDtoEUR)
	case valueFrom == EUR && valueTo == RUB:
		fmt.Printf("Перевод из %s в %s - %.2f", valueFrom, valueTo, value / USDtoEUR * USDtoRUB)
	case valueFrom == RUB && valueTo == EUR:
		fmt.Printf("Перевод из %s в %s - %.2f", valueFrom, valueTo, value * USDtoEUR / USDtoRUB)
	default:
		fmt.Printf("Перевода из %s в %s не найдено", valueFrom, valueTo)
	}
}

func main(){
	var continueMsg string
	fmt.Println("__Конвертер валют__")
	for {
		valueFrom, valueTo, value := userInput();
		calculateValue(valueFrom, valueTo, value)
		fmt.Print("\nПродолжить? да/нет - ")
		fmt.Scan(&continueMsg)

		if continueMsg != "да" {
			break
		}
	}
}