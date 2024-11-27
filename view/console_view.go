package view

import (
	"encoding/json"
	"fmt"
)

type ConsoleView struct{}

func NewConsoleView() *ConsoleView {
	return &ConsoleView{}
}

func (v *ConsoleView) ShowErrorMessage(message string) {
	fmt.Println("\nError:", message)
}

func (v *ConsoleView) ShowWelcomeMessage() {
	fmt.Println("Welcome to the Star Wars API!")
}

func (v *ConsoleView) ShowOptions() int {
	var option int
	fmt.Println("\nPlease choose an option:")
	fmt.Println("1. View Star Wars Films and Planets")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
	fmt.Scanf("%d", &option)
	return option
}

func (v *ConsoleView) ShowFilmsAsJSON(filmPlanets map[string][]string) {
	// Create a map to hold the final JSON structure
	result := make(map[string][]string)

	// Fill the map with sorted films and their corresponding planets
	for title, planets := range filmPlanets {
		result[title] = planets
	}

	// Pretty-print the JSON output
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		return
	}

	// Output the JSON result
	fmt.Println(string(jsonData))
}

func (v *ConsoleView) ShowGoodbyeMessage() {
	fmt.Println("\nThank you for using the Star Wars API viewer. Goodbye!")
}

func (v *ConsoleView) ShowInvalidOptionMessage() {
	fmt.Println("\nInvalid option. Please try again.")
}
