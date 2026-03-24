package formatter

import (
	"encoding/xml"
	"io"
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

type EpisodeNFO struct {
	XMLName  xml.Name `xml:"episodedetails"`
	Title    string   `xml:"title"`
	Season   int      `xml:"season"`
	Episode  int      `xml:"episode"`
	Plot     string   `xml:"plot"`
	Rating   float64  `xml:"rating"`
	Aired    string   `xml:"aired"`
	Thumb    string   `xml:"thumb,omitempty"`
	Runtime  int      `xml:"runtime"`
	UniqueID UniqueID `xml:"uniqueid"`
}

type CollectionNFO struct {
	XMLName xml.Name `xml:"set"`
	Title   string   `xml:"title"`
	Plot    string   `xml:"plot"`
	Thumb   string   `xml:"thumb,omitempty"`
}

type FindResultsNFO struct {
	XMLName xml.Name `xml:"results"`
	Movies  int      `xml:"movies,attr"`
	TV      int      `xml:"tvshows,attr"`
}

func printNFO(w io.Writer, data interface{}, itemType string) error {
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
	case *api.TVEpisode:
		thumb := ""
		if v.StillPath != "" {
			thumb = fmt.Sprintf("https://image.tmdb.org/t/p/w500%s", v.StillPath)
		}
		output = &EpisodeNFO{
			Title:    v.Name,
			Season:   v.SeasonNumber,
			Episode:  v.EpisodeNumber,
			Plot:     v.Overview,
			Rating:   v.VoteAverage,
			Aired:    v.AirDate,
			Thumb:    thumb,
			Runtime:  v.Runtime,
			UniqueID: UniqueID{Type: "tmdb", Default: true, Value: v.ID},
		}
	case *api.Collection:
		thumb := ""
		if v.PosterPath != "" {
			thumb = fmt.Sprintf("https://image.tmdb.org/t/p/w500%s", v.PosterPath)
		}
		output = &CollectionNFO{
			Title: v.Name,
			Plot:  v.Overview,
			Thumb: thumb,
		}
	case *api.FindResults:
		fmt.Fprintln(w, "<!-- Warning: Searching outputs NFO as a generic XML XML. For standard media metadata, fetch specific movie/tv details. -->")
		output = &FindResultsNFO{
			Movies: len(v.MovieResults),
			TV:     len(v.TVResults),
		}
	case *api.SearchResultPage:
	    fmt.Fprintln(w, "<!-- Warning: Searching outputs NFO as a raw XML. For standard media metadata, fetch specific movie/tv details. -->")
		// Fallback to basic struct dump
		output = v
	default:
		output = data
	}

	out, err := xml.MarshalIndent(output, "", "  ")
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n%s\n", string(out))
	return nil
}
