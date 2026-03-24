package formatter

import (
	"encoding/xml"
	"fmt"

	"github.com/gahoolee/tmdb-cli/api"
)

type UniqueID struct {
	Type    string `xml:"type,attr"`
	Default bool   `xml:"default,attr"`
	Value   int    `xml:",chardata"`
}

type MovieNFO struct {
	XMLName       xml.Name `xml:"movie"`
	Title         string   `xml:"title"`
	OriginalTitle string   `xml:"originaltitle"`
	UserRating    float64  `xml:"userrating"`
	Year          string   `xml:"year"`
	Plot          string   `xml:"plot"`
	Tagline       string   `xml:"tagline"`
	Runtime       int      `xml:"runtime"`
	UniqueID      UniqueID `xml:"uniqueid"`
	Thumb         string   `xml:"thumb,omitempty"`
	Genres        []string `xml:"genre"`
	Premiered     string   `xml:"premiered"`
}

type TVNFO struct {
	XMLName    xml.Name `xml:"tvshow"`
	Title      string   `xml:"title"`
	UserRating float64  `xml:"userrating"`
	Year       string   `xml:"year"`
	Plot       string   `xml:"plot"`
	Status     string   `xml:"status"`
	Premiered  string   `xml:"premiered"`
	UniqueID   UniqueID `xml:"uniqueid"`
	Thumb      string   `xml:"thumb,omitempty"`
	Genres     []string `xml:"genre"`
}

type SeasonNFO struct {
	XMLName      xml.Name `xml:"season"`
	Title        string   `xml:"title"`
	SeasonNumber int      `xml:"seasonnumber"`
	Plot         string   `xml:"plot"`
	UniqueID     UniqueID `xml:"uniqueid"`
	Thumb        string   `xml:"thumb,omitempty"`
	Premiered    string   `xml:"premiered"`
}

func printNFO(data interface{}, itemType string) error {
	var output interface{}

	switch v := data.(type) {
	case *api.MovieDetails:
		var genres []string
		for _, g := range v.Genres {
			genres = append(genres, g.Name)
		}
		thumb := ""
		if v.PosterPath != "" {
			thumb = fmt.Sprintf("https://image.tmdb.org/t/p/w500%s", v.PosterPath)
		}
		output = &MovieNFO{
			Title:         v.Title,
			OriginalTitle: v.OriginalTitle,
			UserRating:    v.VoteAverage,
			Year:          extractYear(v.ReleaseDate),
			Plot:          v.Overview,
			Tagline:       v.Tagline,
			Runtime:       v.Runtime,
			UniqueID:      UniqueID{Type: "tmdb", Default: true, Value: v.ID},
			Thumb:         thumb,
			Genres:        genres,
			Premiered:     v.ReleaseDate,
		}
	case *api.TVDetails:
		var genres []string
		for _, g := range v.Genres {
			genres = append(genres, g.Name)
		}
		thumb := ""
		if v.PosterPath != "" {
			thumb = fmt.Sprintf("https://image.tmdb.org/t/p/w500%s", v.PosterPath)
		}
		output = &TVNFO{
			Title:      v.Name,
			UserRating: v.VoteAverage,
			Year:       extractYear(v.FirstAirDate),
			Plot:       v.Overview,
			Status:     v.Status,
			Premiered:  v.FirstAirDate,
			UniqueID:   UniqueID{Type: "tmdb", Default: true, Value: v.ID},
			Thumb:      thumb,
			Genres:     genres,
		}
	case *api.TVSeason:
		thumb := ""
		if v.PosterPath != "" {
			thumb = fmt.Sprintf("https://image.tmdb.org/t/p/w500%s", v.PosterPath)
		}
		output = &SeasonNFO{
			Title:        v.Name,
			SeasonNumber: v.SeasonNumber,
			Plot:         v.Overview,
			UniqueID:     UniqueID{Type: "tmdb", Default: true, Value: v.ID},
			Thumb:        thumb,
			Premiered:    v.AirDate,
		}
	case *api.SearchResultPage:
	    fmt.Println("Warning: Searching outputs NFO as a raw XML. For standard media metadata, fetch specific movie/tv details.")
		// Fallback to basic struct dump
		output = v
	default:
		output = data
	}

	out, err := xml.MarshalIndent(output, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n%s\n", string(out))
	return nil
}
