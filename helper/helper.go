package helper

import "strings"

// Go permite múltiplos retornos, e para cada retorno deve ser especificado o tipo
// o primeiro parênteses são as entradas, o segundo parênteses são os tipos dos retornos

// export: Capitalize first letter
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
