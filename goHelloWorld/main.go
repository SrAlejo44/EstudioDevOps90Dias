package main

import "fmt"

func main() {

	var challenge = "#90DaysOfDevOps"
	const daystotal int = 90
	var daysComplete int
	var twitterName string
	var remainingDays int

	fmt.Println("Enter Your Twitter Handle: ")
	fmt.Scanln(&twitterName)
	fmt.Println("How many days have you completed?: ")
	fmt.Scanln(&daysComplete)

	remainingDays = daystotal - daysComplete

	fmt.Printf("Welcome to %v\n", challenge)
	fmt.Printf("This is a %v challenge\n", daystotal)
	fmt.Printf("Thank you %v for taking part and completing %v days.\n", twitterName, daysComplete)
	fmt.Printf("You have %v days remaining for the %v challenge\n", remainingDays, challenge)
	fmt.Println("Good luck")
}
