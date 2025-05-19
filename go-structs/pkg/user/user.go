package user

import (
	"errors"
	"fmt"
	"time"
)

// User struct only be available to the outer files not its properties.
// bcoz, capital letter rule is applicable to the properties of the struct.
// To access the properties of the struct, change the property first letter to capital letter
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// struct Embeddings
// similar to Inheritance in other programmings
type Admin struct {
	email    string
	password string
	User     // use it directly for anonymous
	// User User // use a property and assign a struct as type. Using caps for exposing it out
}

// adding the (u User) b4 the function name, this function will be converted to method
// that pointing to the specific struct in this case it is User
// (u User) will be the copy of the actual User details
// we can use (u *User) with pointer for avoiding the memory creation,
// but its ok for just reading the details
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// here we have to use the pointer, bcoz we are manipulating the actual data,
// if pointer is not added then the change will be applied to the copy created in the method argument
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(userEmail, userPassword string) Admin {
	return Admin{
		email:    userEmail,
		password: userPassword,
		// User: User{} // Use 'User' as a property for anonymous, if have follow next type
		User: User{
			firstName: "Admin",
			lastName:  "1",
			birthDate: "01/01/0101",
			createdAt: time.Now(),
		},
	}
}

// creating a constructor for the User struct
// it is not in-build with Go, we are creating this builder function
// we are returning the pointer instead of value for avoiding the duplication. its just an option not mandatory
func New(userFirstName, userLastName, userBirthdate string) (*User, error) {

	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("First name, Last name & Birth date is required")
	}

	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}
