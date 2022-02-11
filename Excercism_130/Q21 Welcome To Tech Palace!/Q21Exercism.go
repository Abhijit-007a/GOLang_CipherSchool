package main

import (
	"fmt"
	"regexp"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", border, welcomeMsg, border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	r := regexp.MustCompile(`[*\n]`)
	clean := r.ReplaceAllString(oldMsg, "")
	cleanAndTrim := strings.Trim(clean, " ")
	return cleanAndTrim
}

func main() {
	WelcomeMessage("Judy")
	AddBorder("Welcome!", 10)
}
