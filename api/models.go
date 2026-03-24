package api

// SearchResultPage represents a generic paginated response from TMDB search
type SearchResultPage struct {
	Page         int           `json:"page"`
	Results      []interface{} `json:"results"`
	TotalPages   int           `json:"total_pages"`
	TotalResults int           `json:"total_results"`
}

// Movie represents a TMDB Movie object
type Movie struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	VoteAverage      float64 `json:"vote_average"`
	PosterPath       string  `json:"poster_path"`
	BackdropPath     string  `json:"backdrop_path"`
	MediaType        string  `json:"media_type,omitempty"`
}

// TVShow represents a TMDB TV Series object
type TVShow struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	OriginalName     string  `json:"original_name"`
	Overview         string  `json:"overview"`
	FirstAirDate     string  `json:"first_air_date"`
	VoteAverage      float64 `json:"vote_average"`
	PosterPath       string  `json:"poster_path"`
	BackdropPath     string  `json:"backdrop_path"`
	MediaType        string  `json:"media_type,omitempty"`
}

// MovieDetails extends Movie with more fields
type MovieDetails struct {
	Movie
	Genres   []Genre `json:"genres"`
	Runtime  int     `json:"runtime"`
	Tagline  string  `json:"tagline"`
	Homepage string  `json:"homepage"`
	ImdbID   string  `json:"imdb_id"`
}

// TVDetails extends TVShow with more fields
type TVDetails struct {
	TVShow
	Genres       []Genre `json:"genres"`
	Status       string  `json:"status"`
	Tagline      string  `json:"tagline"`
	Homepage     string  `json:"homepage"`
	NumberOfEps  int     `json:"number_of_episodes"`
	NumberOfSeas int     `json:"number_of_seasons"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TVSeason represents a TMDB TV Season object
type TVSeason struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	AirDate      string  `json:"air_date"`
	PosterPath   string  `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	VoteAverage  float64 `json:"vote_average"`
}
