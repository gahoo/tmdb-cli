package formatter

import (
	"encoding/json"
	"io"
	"fmt"
	"strings"

	"github.com/gahoolee/tmdb-cli/api"
)

func printMarkdown(w io.Writer, data interface{}, itemType string) error {
	switch v := data.(type) {
	case *api.SearchResultPage:
		return printSearchMarkdown(w, v, itemType)
	case *api.MovieDetails:
		return printMovieMarkdown(w, v)
	case *api.TVDetails:
		return printTVMarkdown(w, v)
	case *api.TVEpisode:
		return printEpisodeMarkdown(w, v)
	case *api.Collection:
		return printCollectionMarkdown(w, v)
	case *api.FindResults:
		return printFindMarkdown(w, v)
	default:
		// Fallback
		out, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "### %s Results\n\n```json\n%s\n```\n", strings.ToUpper(itemType), string(out))
		return nil
	}
}

func printSearchMarkdown(w io.Writer, data *api.SearchResultPage, itemType string) error {
	fmt.Fprintf(w, "# Search Results (Type: %s, Total: %d, Page %d/%d)\n\n", strings.ToUpper(itemType), data.TotalResults, data.Page, data.TotalPages)
	for idx, res := range data.Results {
		// results elements are interface{}, map to json to extract fields safely
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

		fmt.Fprintf(w, "%d. **%s** (%s) - Type: %s | ID: %.0f\n", idx+1, title, date, mediaType, id)
	}
	return nil
}

func printMovieMarkdown(w io.Writer, data *api.MovieDetails) error {
	fmt.Fprintf(w, "# 🎬 Movie: %s (%s)\n\n", data.Title, extractYear(data.ReleaseDate))
	if data.Tagline != "" {
		fmt.Fprintf(w, "> *%s*\n\n", data.Tagline)
	}
	fmt.Fprintf(w, "**Original Title:** %s\n", data.OriginalTitle)
	fmt.Fprintf(w, "**Rating:** %.1f\n", data.VoteAverage)
	fmt.Fprintf(w, "**Runtime:** %d minutes\n", data.Runtime)
	
	if len(data.Genres) > 0 {
		var genres []string
		for _, g := range data.Genres {
			genres = append(genres, g.Name)
		}
		fmt.Fprintf(w, "**Genres:** %s\n", strings.Join(genres, ", "))
	}
	fmt.Fprintf(w, "\n### Overview\n%s\n", data.Overview)
	
	if data.PosterPath != "" {
		fmt.Fprintf(w, "\n### Poster\n`https://image.tmdb.org/t/p/w500%s`\n", data.PosterPath)
	}
	return nil
}

func printTVMarkdown(w io.Writer, data *api.TVDetails) error {
	fmt.Fprintf(w, "# 📺 TV: %s (%s)\n\n", data.Name, extractYear(data.FirstAirDate))
	if data.Tagline != "" {
		fmt.Fprintf(w, "> *%s*\n\n", data.Tagline)
	}
	fmt.Fprintf(w, "**Status:** %s\n", data.Status)
	fmt.Fprintf(w, "**Rating:** %.1f\n", data.VoteAverage)
	fmt.Fprintf(w, "**Seasons:** %d | **Episodes:** %d\n", data.NumberOfSeas, data.NumberOfEps)
	
	if len(data.Genres) > 0 {
		var genres []string
		for _, g := range data.Genres {
			genres = append(genres, g.Name)
		}
		fmt.Fprintf(w, "**Genres:** %s\n", strings.Join(genres, ", "))
	}
	fmt.Fprintf(w, "\n### Overview\n%s\n", data.Overview)
	return nil
}

func extractYear(date string) string {
	if len(date) >= 4 {
		return date[:4]
	}
	return "N/A"
}

func printEpisodeMarkdown(w io.Writer, data *api.TVEpisode) error {
	fmt.Fprintf(w, "# 📺 TV Episode: %s (S%02dE%02d)\n\n", data.Name, data.SeasonNumber, data.EpisodeNumber)
	fmt.Fprintf(w, "**Air Date:** %s\n", data.AirDate)
	fmt.Fprintf(w, "**Rating:** %.1f\n", data.VoteAverage)
	fmt.Fprintf(w, "\n### Overview\n%s\n", data.Overview)
	
	if len(data.GuestStars) > 0 {
		var stars []string
		for _, person := range data.GuestStars {
			stars = append(stars, fmt.Sprintf("%s as %s", person.Name, person.Character))
		}
		fmt.Fprintf(w, "\n### Guest Stars\n- %s\n", strings.Join(stars, "\n- "))
	}
	return nil
}

func printCollectionMarkdown(w io.Writer, data *api.Collection) error {
	fmt.Fprintf(w, "# 🎬 Collection: %s\n\n", data.Name)
	fmt.Fprintf(w, "### Overview\n%s\n", data.Overview)
	
	if len(data.Parts) > 0 {
		fmt.Fprintf(w, "\n### Movies in Collection\n")
		for idx, part := range data.Parts {
			fmt.Fprintf(w, "%d. %s (%s)\n", idx+1, part.Title, extractYear(part.ReleaseDate))
		}
	}
	return nil
}

func printFindMarkdown(w io.Writer, data *api.FindResults) error {
	fmt.Fprintf(w, "# 🔍 Find Results\n\n")
	idx := 1
	for _, m := range data.MovieResults {
		fmt.Fprintf(w, "%d. 🎬 Movie: **%s** (%s) - ID: %d\n", idx, m.Title, extractYear(m.ReleaseDate), m.ID)
		idx++
	}
	for _, t := range data.TVResults {
		fmt.Fprintf(w, "%d. 📺 TV: **%s** (%s) - ID: %d\n", idx, t.Name, extractYear(t.FirstAirDate), t.ID)
		idx++
	}
	for _, p := range data.PersonResults {
		fmt.Fprintf(w, "%d. 👤 Person: **%s** - ID: %d\n", idx, p.Name, p.ID)
		idx++
	}
	for _, e := range data.TVEpisodeResults {
		fmt.Fprintf(w, "%d. 🎞️ Episode: **%s** (S%02dE%02d) - ID: %d\n", idx, e.Name, e.SeasonNumber, e.EpisodeNumber, e.ID)
		idx++
	}
	for _, s := range data.TVSeasonResults {
		fmt.Fprintf(w, "%d. 📦 Season: **%s** - ID: %d\n", idx, s.Name, s.ID)
		idx++
	}
	if idx == 1 {
		fmt.Fprintf(w, "*No results found*\n")
	}
	return nil
}

func printDynamicMarkdown(w io.Writer, data interface{}) error {
	printMDNode(w, data, 0)
	return nil
}

func printMDNode(w io.Writer, node interface{}, depth int) {
	indent := strings.Repeat("  ", depth)
	switch v := node.(type) {
	case map[string]interface{}:
		for key, val := range v {
			if isComplex(val) {
				fmt.Fprintf(w, "%s**%s**:\n", indent, key)
				printMDNode(w, val, depth+1)
			} else {
				fmt.Fprintf(w, "%s**%s**: %v\n", indent, key, val)
			}
		}
	case []interface{}:
		for idx, item := range v {
			if isComplex(item) {
				fmt.Fprintf(w, "%s- Item %d:\n", indent, idx+1)
				printMDNode(w, item, depth+1)
			} else {
				fmt.Fprintf(w, "%s- %v\n", indent, item)
			}
		}
	default:
		fmt.Fprintf(w, "%s%v\n", indent, v)
	}
}

func isComplex(node interface{}) bool {
	switch node.(type) {
	case map[string]interface{}, []interface{}:
		return true
	}
	return false
}
