package formatter

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/gahoolee/tmdb-cli/api"
)

func printTable(data interface{}, itemType string) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)

	switch v := data.(type) {
	case *api.SearchResultPage:
		fmt.Fprintln(w, "ID\tTITLE\tDATE\tTYPE\tVOTE")
		fmt.Fprintln(w, "--\t-----\t----\t----\t----")
		for _, res := range v.Results {
			b, _ := json.Marshal(res)
			var temp map[string]interface{}
			json.Unmarshal(b, &temp)

			var title, date, mediaType string
			if temp["title"] != nil {
				title = temp["title"].(string)
			} else if temp["name"] != nil {
				title = temp["name"].(string)
			}
			
			if temp["release_date"] != nil {
				date = temp["release_date"].(string)
			} else if temp["first_air_date"] != nil {
				date = temp["first_air_date"].(string)
			}

			if temp["media_type"] != nil {
				mediaType = temp["media_type"].(string)
			} else {
				mediaType = itemType
			}

			id := temp["id"].(float64)
			vote := temp["vote_average"].(float64)

			fmt.Fprintf(w, "%.0f\t%s\t%s\t%s\t%.1f\n", id, title, extractYear(date), mediaType, vote)
		}

	case *api.MovieDetails:
		fmt.Fprintln(w, "KEY\tVALUE")
		fmt.Fprintln(w, "---\t-----")
		fmt.Fprintf(w, "ID\t%d\n", v.ID)
		fmt.Fprintf(w, "Title\t%s\n", v.Title)
		fmt.Fprintf(w, "Original\t%s\n", v.OriginalTitle)
		fmt.Fprintf(w, "Year\t%s\n", extractYear(v.ReleaseDate))
		fmt.Fprintf(w, "Rating\t%.1f\n", v.VoteAverage)
		fmt.Fprintf(w, "Runtime\t%d min\n", v.Runtime)
		
		var genres []string
		for _, g := range v.Genres {
			genres = append(genres, g.Name)
		}
		fmt.Fprintf(w, "Genres\t%s\n", strings.Join(genres, ", "))

	case *api.TVDetails:
		fmt.Fprintln(w, "KEY\tVALUE")
		fmt.Fprintln(w, "---\t-----")
		fmt.Fprintf(w, "ID\t%d\n", v.ID)
		fmt.Fprintf(w, "Name\t%s\n", v.Name)
		fmt.Fprintf(w, "Original\t%s\n", v.OriginalName)
		fmt.Fprintf(w, "First Aired\t%s\n", extractYear(v.FirstAirDate))
		fmt.Fprintf(w, "Rating\t%.1f\n", v.VoteAverage)
		fmt.Fprintf(w, "Seasons\t%d\n", v.NumberOfSeas)
		fmt.Fprintf(w, "Episodes\t%d\n", v.NumberOfEps)
		
		var genres []string
		for _, g := range v.Genres {
			genres = append(genres, g.Name)
		}
		fmt.Fprintf(w, "Genres\t%s\n", strings.Join(genres, ", "))

	default:
		fmt.Fprintln(w, "Unsupported table view for this data type.")
	}

	w.Flush()
	return nil
}
