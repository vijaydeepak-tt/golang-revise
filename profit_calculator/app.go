package main

import (
	"errors"
	"fmt"
	"os"
)

func saveResult(ebt, profit, ratio float64) {
	data := fmt.Sprintf("EBT: %0.0f\nProfit: %0.2f\nRatio: %0.3f\n", ebt, profit, ratio)
	os.WriteFile("income.txt", []byte(data), 0644)
}

func main() {
	revenue, err := getUserInputs("Revenue: ")
	fireError(err)
	expense, err := getUserInputs("Expense: ")
	fireError(err)
	tax, err := getUserInputs("Tax: ")
	fireError(err)

	ebt, profit, ratio := calculateFinancials(revenue, expense, tax)
	saveResult(ebt, profit, ratio)
	fmt.Printf("EBT: %0.0f  Profit: %0.2f  Ratio: %0.3f\n", ebt, profit, ratio)

}

func fireError(err error) {
	if err != nil {
		fmt.Println(err)
		// panic(err)
		return
	}
}

func getUserInputs(label string) (float64, error) {
	var input float64
	fmt.Print(label)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("Invalid input. Value must be greater than 0.")
	}

	return input, nil
}

func calculateFinancials(revenue, expense, tax float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expense
	profit = ebt * (1 - tax/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
