package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName string = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your email: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("The whole Slice %v\n", bookings)
		// fmt.Printf("The first element in the Slice %v\n", bookings[0])
		// fmt.Printf("The type of the Slice %T\n", bookings)
		// fmt.Printf("The size of the Slice %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of the bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("Our %v is booked out. Come back next year.\n", conferenceName)
			break
		}
	}

}
