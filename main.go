package main

import "fmt"

func main() {
	const eventName string = "Party"
	
	// read more about different ways % is used in formatting string here: https://golang.org/pkg/fmt/
	fmt.Printf("Welcome to the %T", eventName)

}