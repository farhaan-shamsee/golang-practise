package main

import "fmt"

const LoginToken string = "ghjkl" // public variable

func main() {
	var username string = "farhaan"
	fmt.Println("Hello, " + username + "!")
	fmt.Printf("variable is of type: %T \n", username)
	
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "farhaan.com"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n", website)

	// no var style
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)
	fmt.Printf("variable is of type: %T \n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}

