package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var eventMaster = "tix"
var eventName = "NBA Finals";
const eventTickets int = 50;
var remainingTickets uint = 50
var bookings []string

func main () {

	greetUsers()
	
	for remainingTickets > 0 && len(bookings) <= 50 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, uint(userTickets), remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			fmt.Printf("We only have %v tickets available, you can't book tickets %v\n", remainingTickets, userTickets)
			continue
		}

		if userTickets <= int(remainingTickets) {

			bookTickets(uint(userTickets), firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of each bookings are: %v\n\n", firstNames)
			
			if remainingTickets == 0 {
				fmt.Println("Our event is booked out. Check back soon!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name input")
			}
			if !isValidEmail {
				fmt.Println("Invalid email address")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticketnumber input")
			}
			fmt.Printf("Your input data is invalid. Please, try again\n")
		}
	}
}

func greetUsers () {
	fmt.Printf("Welcome to %v, your preferred event ticketing platform for memorable experiences\n", eventMaster)
	fmt.Printf("Get tickets to your favorite events such as %v\n", eventName);
	fmt.Printf("We have total of %v tickets available for the show and %v left\n", eventTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput () (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("How many tickets do you want? ")
	fmt.Scan(&userTickets)

	fmt.Println("Enter your email address ")
	fmt.Scan(&email)

	return firstName, lastName, email, userTickets;
}

func bookTickets( userTickets uint, firstName string, lastName string, email string ) (uint, []string) {
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you for using Tix for your ticketing, %v %v.\nYou have just bought %v tickets.\nA confirmation message will be sent to %v shortly\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets left for the %v\n", remainingTickets, eventName)

	return remainingTickets, bookings
}