package films

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type FilmService struct{}

func NewFilmService() *FilmService {
	return &FilmService{}
}

func (s *FilmService) FetchFilms() ([]Film, error) {
	resp, err := http.Get("https://swapi.dev/api/films/")
	if err != nil {
		return nil, fmt.Errorf("Error fetching films: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Results []Film `json:"results"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Error decoding response: %v", err)
	}

	return result.Results, nil
}

func (s *FilmService) SortFilmsByReleaseDate(films []Film) []string {
	sort.Slice(films, func(i, j int) bool {
		return films[i].ReleaseDate < films[j].ReleaseDate
	})

	titles := make([]string, len(films))
	for i, film := range films {
		titles[i] = film.Title
	}
	return titles
}
