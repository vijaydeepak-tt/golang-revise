package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var accountFile = "account.txt"

func getBalance() (float64, error) {
	data, err := os.ReadFile(accountFile)

	if err != nil {
		return 0, errors.New("Failed to find Account file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("Failed to parse the balance from file.")
	}

	return balance, nil
}

func saveBalance(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountFile, []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = getBalance()

	if err != nil {
		fmt.Println("ERROR!!!")
		fmt.Println(err)
		fmt.Println("=============================")
		// return
		// panic("Can't continue further, reach to the Go Bank!")
	}

	fmt.Println("Welcome to the Go Bank!")

	// withIfElse(accountBalance)
	withSwitch(accountBalance)
}

func getUserInputs(label string) float64 {
	var input float64
	fmt.Print(label)
	fmt.Scan(&input)

	return input
}

func withIfElse(accountBalance float64) {
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your Account balance: ", accountBalance)
		} else if choice == 2 {
			depositAmount := getUserInputs("Your Deposit Amount: ")

			if depositAmount <= 0 {
				fmt.Print("Invalid Amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your New balance: ", accountBalance)
		} else if choice == 3 {
			withdrawAmount := getUserInputs("Your Withdraw Amount: ")

			if withdrawAmount <= 0 {
				fmt.Print("Invalid Amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Print("Invalid Amount. You can't withdraw amount more than you have.")
				fmt.Println("Your Account balance: ", accountBalance)
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your New balance: ", accountBalance)
		} else {
			fmt.Print("Goodbye!!!")
			break
		}
	}
	fmt.Print("Thank you for Banking with Go!!")
}

func withSwitch(accountBalance float64) {
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Account balance: ", accountBalance)
		case 2:
			depositAmount := getUserInputs("Your Deposit Amount: ")

			if depositAmount <= 0 {
				fmt.Print("Invalid Amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			saveBalance(accountBalance)
			fmt.Println("Your New balance: ", accountBalance)
		case 3:
			withdrawAmount := getUserInputs("Your Withdraw Amount: ")

			if withdrawAmount <= 0 {
				fmt.Print("Invalid Amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Print("Invalid Amount. You can't withdraw amount more than you have.")
				fmt.Println("Your Account balance: ", accountBalance)
				continue
			}

			accountBalance -= withdrawAmount
			saveBalance(accountBalance)
			fmt.Println("Your New balance: ", accountBalance)
		default:
			fmt.Println("Goodbye!!!")
			fmt.Println("Thank you for Banking with Go!!")
			return
			// break
		}

	}

}
