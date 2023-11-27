package main

import (
	"fmt"
	"strings"
)

func main () {

	eventMaster := "tix"
	var eventName = "NBA Finals";
	const eventTickets int = 50;
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v, your preferred event ticketing platform for memorable experiences\n", eventMaster);
	fmt.Printf("We have total of %v tickets available for the show and %v left\n", eventTickets, remainingTickets)
	fmt.Printf("Get tickets to your favorite events such as %v\n", eventName);

	for remainingTickets > 0 && len(bookings) <= 50 {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < int(remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			fmt.Printf("We only have %v tickets available, you can't book tickets %v\n", remainingTickets, userTickets)
			continue
		}

		if userTickets <= int(remainingTickets) {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName + " " + lastName)


			fmt.Printf("Thank you for using Tix for your ticketing, %v %v.\nYou have just bought %v tickets.\nA confirmation message will be sent to %v shortly\n", firstName, lastName, userTickets, email)
			fmt.Printf("We have %v tickets left for the %v\n", remainingTickets, eventName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("First names of each booking: %v \n", firstNames)
			
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