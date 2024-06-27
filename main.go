package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	bookings := []string{}
	firstNames := []string{} // Moved outside the loop

	fmt.Printf("Welcome to %v booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your Tickets here to attend!")

	for {
		var firstName string
		var email string
		var userTickets uint
		var lastName string

		fmt.Println("Enter your first Name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last Name") // Corrected typo: "first Last Name" -> "last Name"
		fmt.Scan(&lastName)

		fmt.Println("Enter your E-Mail Address") // Corrected typo: "Adress" -> "Address"
		fmt.Scan(&email)

		fmt.Println("Enter number of Tickets")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v Tickets remaining. You can not book %v Tickets!\n", remainingTickets, userTickets)
			continue
		}

		if userTickets <= remainingTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v Tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v Tickets remaining for %v.\n", remainingTickets, conferenceName)

			// Extract first names
			names := strings.Fields(firstName + " " + lastName)
			if len(names) > 0 {
				firstNames = append(firstNames, names[0])
			}
		} else {
			fmt.Printf("We only have %v Tickets remaining. You can't book %v Tickets.\n", remainingTickets, userTickets)
		}

		fmt.Printf("The first Names of Bookings are: %v\n", firstNames)
		fmt.Printf("These are all our Bookings: %v\n", bookings)

		if remainingTickets == 0 {
			fmt.Println("Our Conference is booked Out! Come back next Year.")
			break
		}

	}
}
