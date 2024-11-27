package controller

import (
	"log"
	"sort"
	"startwars-api/films"
	"startwars-api/planets"
	"startwars-api/view"
)

type StarWarsController struct {
	FilmService   *films.FilmService
	PlanetService *planets.PlanetService
	View          *view.ConsoleView
}

func NewStarWarsController(filmService *films.FilmService, planetService *planets.PlanetService, view *view.ConsoleView) *StarWarsController {
	return &StarWarsController{
		FilmService:   filmService,
		PlanetService: planetService,
		View:          view,
	}
}

func (c *StarWarsController) Run() {
	// Display welcome message
	c.View.ShowWelcomeMessage()

	// Display options to the user
	for {
		option := c.View.ShowOptions()
		switch option {
		case 1:
			// Fetch films when user selects this option
			films, err := c.FilmService.FetchFilms()
			if err != nil {
				log.Printf("Error fetching films: %v", err)
				c.View.ShowErrorMessage("Error fetching films, please try again later.")
				continue
			}

			// Sort films by release date
			sort.Slice(films, func(i, j int) bool {
				return films[i].ReleaseDate < films[j].ReleaseDate
			})

			// Build film-to-planets mapping
			filmPlanets := make(map[string][]string)
			for _, film := range films {
				planetNames, err := c.PlanetService.FetchPlanetNames(film.PlanetURLs)
				if err != nil {
					log.Printf("Error fetching planets for film %s: %v", film.Title, err)
					c.View.ShowErrorMessage("Error fetching planets, please try again later.")
					continue
				}
				filmPlanets[film.Title] = planetNames
			}

			// Show films and planets in JSON format
			c.View.ShowFilmsAsJSON(filmPlanets)

		case 0:
			// Exit the program
			c.View.ShowGoodbyeMessage()
			return

		default:
			// Handle invalid options
			c.View.ShowInvalidOptionMessage()
		}
	}
}
