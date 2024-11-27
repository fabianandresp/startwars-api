package main

import (
	"startwars-api/controller"
	"startwars-api/films"
	"startwars-api/planets"
	"startwars-api/view"
)

func main() {
	// Build instance of services
	filmService := films.NewFilmService()
	planetService := planets.NewPlanetService()
	consoleView := view.NewConsoleView()

	// Build the controller
	starWarsController := controller.NewStarWarsController(filmService, planetService, consoleView)

	// Controllet execute
	starWarsController.Run()
}
