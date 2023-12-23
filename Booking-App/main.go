package main

import (
	"fmt"
	"strings"
)

// need entrypoint into application - go needs to know where to start executing code
func main() { //starts executing here
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	for {
		//slice, dynamic in size unlike array
		//or can do bookings = []string{}
		//or bookings := []string{}

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}           //range iterates over elements for different data structures (not just array and slices), for arrays and slices, range provides index and value
		for _, booking := range bookings { //      _ is used to ignore variable (index), so won't give error
			var names = strings.Fields(booking) //returns splice
			firstNames = append(firstNames, names[0])

		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		//noTicketsRemaining  := remainingTickets == 0

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}

	}

}
