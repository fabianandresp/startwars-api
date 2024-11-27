package planets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type PlanetService struct{}

func NewPlanetService() *PlanetService {
	return &PlanetService{}
}

func (s *PlanetService) FetchPlanetNames(urls []string) ([]string, error) {
	var planetNames []string
	for _, url := range urls {
		// Get planet data from the API using its URL
		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("Error fetching planet: %v", err)
		}
		defer resp.Body.Close()

		var planet struct {
			Name string `json:"name"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&planet); err != nil {
			return nil, fmt.Errorf("Error decoding planet response: %v", err)
		}

		planetNames = append(planetNames, planet.Name)
	}

	// Sort planet names alphabetically
	sort.Strings(planetNames)

	return planetNames, nil
}
