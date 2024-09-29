package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName string = "Go conference" //for package level variables we cannot declare varibles like conferenceName:="Go conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0) //list of map
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetuser()

	for {

		//call userInput
		firstName, lastName, email, userTickets := userInput()

		//call validatUserInput
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//call bookingTicket
			bookTicket(firstName, lastName, email, userTickets)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email) //to make the code concurrent go->starts a new goroutine

			// fmt.Printf("the whole slice :%v\n", bookings)
			// fmt.Printf("the first value: %v\n", bookings[0])
			// fmt.Printf("the slice type: %T\n", bookings)
			// fmt.Printf("the slice length %v\n", len(bookings))

			//call function getFirstNames
			firstnames := getFirstNames()
			fmt.Printf("the first names of bookings are:%v\n", firstnames)

			if remainingTickets == 0 {
				fmt.Println("our conference is booked out. Come eback next year")
				break

			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short. Try again")
			}
			if !isValidEmail {
				fmt.Println("email address you entered does not contain @ sign. Try again")
			}
			if !isValidTicketNumber {
				fmt.Println("You have entered an invalid ticket number. Try again")
			}

		}

	}
	wg.Wait()
}

func greetuser() {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("We have toatal of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName) //firstNames = append(firstNames, booking["firstName"]) for maps
	}
	return firstNames
}

func userInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint
	var email string

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter the the number of tickets you want to book")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	//var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve an email conformation on %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("The number of remaining tickets are %v\n", remainingTickets)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########################")
	fmt.Printf("Sending ticket:%v\n to email adderss %v\n", ticket, email)
	fmt.Println("##########################")
	wg.Done()
}

//switch working
// 	var city string
// 	fmt.Println("Enter city")
// 	fmt.Scan(&city)
// 	switch city {
// 	case "Bangalore":
// 		fmt.Println("city is bangalore")

// 	case "Mangalore":
// 		fmt.Println("city is mangalore")

// 	case "Mysore":
// 		fmt.Println("city is Mysore")

// 	default:
// 		fmt.Println("Inavlid city")
// 	}

// }
