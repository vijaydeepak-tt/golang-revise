package main

import (
	"fmt"

	"example.com/bank/pkg/fileops"
	"github.com/Pallinder/go-randomdata"
)

var accountFile = "account.txt"

func main() {

	var accountBalance, err = fileops.GetFloatValue(accountFile)

	if err != nil {
		fmt.Println("ERROR!!!")
		fmt.Println(err)
		fmt.Println("=============================")
		// return
		// panic("Can't continue further, reach to the Go Bank!")
	}

	fmt.Println("Welcome to the Go Bank!")
	fmt.Println("Reach us 24/7 at ", randomdata.PhoneNumber())

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
		presentOptions()

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
		presentOptions()

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
			fileops.SaveFloatValue(accountBalance, accountFile)
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
			fileops.SaveFloatValue(accountBalance, accountFile)
			fmt.Println("Your New balance: ", accountBalance)
		default:
			fmt.Println("Goodbye!!!")
			fmt.Println("Thank you for Banking with Go!!")
			return
			// break
		}

	}

}
