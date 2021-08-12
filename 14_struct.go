package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthdate string
	createdDate time.Time
}

var reader2 = bufio.NewReader(os.Stdin)

func main()  {
	var newUser User
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	created := time.Now()
	newUser = User{firstName: firstName, lastName: lastName, birthdate: birthdate, createdDate: created}
	outputUserDetails(newUser)
}
func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader2.ReadString('\n')
	cleanedInput := strings.Replace(userInput, "\n", "", -1)
	return cleanedInput
}
func outputUserDetails(user User)  {
	fmt.Printf("My name is %v %v (born on %v)", user.firstName, user.lastName, user.birthdate)
}