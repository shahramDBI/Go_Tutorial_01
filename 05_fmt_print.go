package main

import "fmt"

func main()  {
	firsName := "Shahram"
	lastName := "DBI"
	fullName := firsName + " " + lastName
	fmt.Println("Hi, my full name is", fullName)
	fmt.Printf("Hi, my full name is %v", fullName)
}
