// C:\Users\edgar\OneDrive\Documents\GO\src\booking-app> go run .

package main

import (
	"fmt"
	"sync"
	"time"

	"booking-app/helper"
)

// package level variables
// a função tem o objetivo de deixar o código mais bonito, compacto
// mas se com isso as funções receberem muitos parâmetros, é conveniente declarar as variáveis globalmente
// fora das funções (inclusive main)
// elas ficam acessíveis a todas as funções e não necessitam ser passadas como parâmetro
const conferenceTickets int = 50

var (
	conferenceName        = "GO Conference"
	remainingTickets uint = 50
	bookings              = make([]UserData, 0)

// bookings              = make([]map[string]string, 0)
)

// type statement - custom types
// type keyword creates a new type
// "create a type called 'UserData' based  on a struct of firstName, lastName..."
// in fact, you could also create a type based on every other data type like string, int etc
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	// for {

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		// var noTicketsRemaining bool = remainingTickets == 0 //boolean type
		if remainingTickets == 0 {
			// program end
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First Name or Last Name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign.")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid.")
		}
	}
	wg.Wait()
	//}

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T and conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	/*fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))*/
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// para cada parâmetro de input, deve ser especificado o seu tipo
// se a função tiver return, deve ser especificado o tipo da função
func getFirstNames() []string {
	// revisar esse conceito
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your e-mail address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a struct for a user
	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// map
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at: %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// função que simula uma tarefa onerosa que demora mto tempo (ex: gerar pdf, enviar por email p/ usuário etc)
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
