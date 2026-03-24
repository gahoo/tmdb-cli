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
	Genres            []Genre            `json:"genres"`
	Runtime           int                `json:"runtime"`
	Tagline           string             `json:"tagline"`
	Homepage          string             `json:"homepage"`
	ImdbID            string             `json:"imdb_id"`
	Credits           *Credits           `json:"credits,omitempty"`
	AlternativeTitles *AlternativeTitles `json:"alternative_titles,omitempty"`
	ExternalIDs       *ExternalIDs       `json:"external_ids,omitempty"`
}

// TVDetails extends TVShow with more fields
type TVDetails struct {
	TVShow
	Genres            []Genre            `json:"genres"`
	Status            string             `json:"status"`
	Tagline           string             `json:"tagline"`
	Homepage          string             `json:"homepage"`
	NumberOfEps       int                `json:"number_of_episodes"`
	NumberOfSeas      int                `json:"number_of_seasons"`
	Credits           *Credits           `json:"credits,omitempty"`
	AlternativeTitles *AlternativeTitles `json:"alternative_titles,omitempty"`
	ExternalIDs       *ExternalIDs       `json:"external_ids,omitempty"`
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

type Credits struct {
	Cast []Cast `json:"cast"`
	Crew []Crew `json:"crew"`
}

type Cast struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Character   string `json:"character"`
	ProfilePath string `json:"profile_path"`
	Order       int    `json:"order"`
}

type Crew struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Job         string `json:"job"`
	Department  string `json:"department"`
	ProfilePath string `json:"profile_path"`
}

type ExternalIDs struct {
	IMDbID      string `json:"imdb_id"`
	TVDBID      int    `json:"tvdb_id"`
	FreebaseID  string `json:"freebase_id"`
	FreebaseMID string `json:"freebase_mid"`
	TVRageID    int    `json:"tvrage_id"`
	WikidataID  string `json:"wikidata_id"`
	FacebookID  string `json:"facebook_id"`
	InstagramID string `json:"instagram_id"`
	TwitterID   string `json:"twitter_id"`
}

type AlternativeTitles struct {
	Titles  []AlternativeTitle `json:"titles,omitempty"`
	Results []AlternativeTitle `json:"results,omitempty"`
}

type AlternativeTitle struct {
	Iso31661 string `json:"iso_3166_1"`
	Title    string `json:"title"`
	Type     string `json:"type"`
}

type TVEpisode struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Overview      string  `json:"overview"`
	AirDate       string  `json:"air_date"`
	EpisodeNumber int     `json:"episode_number"`
	SeasonNumber  int     `json:"season_number"`
	VoteAverage   float64 `json:"vote_average"`
	StillPath     string  `json:"still_path"`
	Runtime       int     `json:"runtime"`
	GuestStars    []Cast  `json:"guest_stars"`
	Crew          []Crew  `json:"crew"`
}

type Collection struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	BackdropPath string  `json:"backdrop_path"`
	Parts        []Movie `json:"parts"`
}

type Configuration struct {
	Images struct {
		BaseURL       string   `json:"base_url"`
		SecureBaseURL string   `json:"secure_base_url"`
		BackdropSizes []string `json:"backdrop_sizes"`
		LogoSizes     []string `json:"logo_sizes"`
		PosterSizes   []string `json:"poster_sizes"`
		ProfileSizes  []string `json:"profile_sizes"`
		StillSizes    []string `json:"still_sizes"`
	} `json:"images"`
}

type FindResults struct {
	MovieResults     []Movie     `json:"movie_results"`
	TVResults        []TVShow    `json:"tv_results"`
	PersonResults    []Person    `json:"person_results"`
	TVEpisodeResults []TVEpisode `json:"tv_episode_results"`
	TVSeasonResults  []TVSeason  `json:"tv_season_results"`
}

type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ProfilePath string `json:"profile_path"`
}
