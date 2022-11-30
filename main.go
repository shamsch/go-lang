package main

import "fmt"

func main() {
	// constant and variable declaration
	const eventName string = "Party"
	const totalGuests uint = 100
	const ticketPrice float32 = 12.55 
	ticketLeft := totalGuests

	// printing welcome message
	// read more about different ways % is used in formatting string here: https://golang.org/pkg/fmt/
	fmt.Printf("Welcome to the %v\n", eventName)
	fmt.Printf("Total guests allowed: %v\n", totalGuests)
	fmt.Printf("Ticket price: $%.fv\n", ticketPrice)
	fmt.Printf("Tickets left: %v\n", ticketLeft)

	// user name and number of tickets
	var userName string
	var tickets uint

	// taking user input
	// variable value is scanned using the pointer of the variable (&)
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&tickets)

	// printing user name and number of tickets
	fmt.Printf("Hi %v, you have booked %v tickets\n", userName, tickets)
}