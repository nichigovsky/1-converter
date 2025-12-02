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

type currencyType = map[string]float64

func calculateValue(valueFrom string, valueTo string, value float64, currency *map[string]currencyType) {
	cur := *currency

	fmt.Printf("Перевод из %s в %s - %.2f", valueFrom, valueTo, value * cur[valueFrom][valueTo])
}

func main(){
	var continueMsg string
	currency := map[string]currencyType{
		"RUB": {
			"USD": 0.013,
			"EUR": 0.011,
		},
		"EUR": {
			"USD": 1.16,
			"RUB": 90.14,
		},
		"USD": {
			"EUR": 0.86,
			"RUB": 77.72,
		},
	}

	fmt.Println("__Конвертер валют__")
	for {
		valueFrom, valueTo, value := userInput();
		calculateValue(valueFrom, valueTo, value, &currency)
		fmt.Print("\nПродолжить? да/нет - ")
		fmt.Scan(&continueMsg)

		if continueMsg != "да" {
			break
		}
	}
}