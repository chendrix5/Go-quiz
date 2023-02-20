package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game! \n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("You can play!")
	} else {
		fmt.Println("Come back when you're older")
		return
	}

	fmt.Printf("What language is this program written in?")
	var answer string
	fmt.Scan(&answer)

	fmt.Println(answer)
		
}

