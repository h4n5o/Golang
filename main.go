package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	// var bookings = [50]string{} !!!!Alternative Schreibweise dafür ist in Zeile 17!!!!
	bookings := []string{}
	// var bookings []string !!!!Alternative Schreibweise dafür ist in Zeile 11
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your Tickets here to attend!")

	for {

		var firstName string
		var email string
		var userTickets uint
		var lastName string
		// User nach Informationen fragen (User Input)
		fmt.Println("Enter your first Name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your first Last Name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your E-Mail Adress")
		fmt.Scan(&email)

		fmt.Println("Enter number of Tickets")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		//Userinformationen speichern
		bookings = append(bookings, firstName+" "+lastName)

		// Gibt Slice Informationen aus !!!!Slices und Arrays!!!!
		fmt.Printf("The whole Slice: %v\n", bookings)
		fmt.Printf("The first Value: %v\n", bookings[0])
		fmt.Printf("The Type of Slice: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v Tickets. You will receive a confirmation Mail to %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v Tickets remaining for %v.\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings { //Underscore ist für Unused Variables
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first Names of Bookings are: %v\n", firstNames)
		fmt.Printf("These are all our Bookings: %v\n", bookings)

	}
}
