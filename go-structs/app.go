package main

import (
	"fmt"

	"example.com/go-structs/pkg/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	var appUser *user.User
	var err error

	appUser, err = user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	adminUser := user.NewAdmin("test@email.com", "***************")

	// Used with defined property names
	// adminUser.User.OutputUserDetails()
	// adminUser.User.ClearUserName()
	// adminUser.User.OutputUserDetails()

	// Used with anonymous struct
	adminUser.OutputUserDetails()
	adminUser.ClearUserName()
	adminUser.OutputUserDetails()

	// outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

// func outputUserDetails(u *user) {
// 	// In structs we no need to depointer to get the value in go
// 	// we can also use (*u).firstName to get value, but u.firstName is the shorter hand
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// using ln will consider the enter will move to next step else it will wait for the value
	fmt.Scanln(&value)
	return value
}
