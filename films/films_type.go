package films

type Film struct {
	Title       string   `json:"title"`
	ReleaseDate string   `json:"release_date"`
	PlanetURLs  []string `json:"planets"`
}
