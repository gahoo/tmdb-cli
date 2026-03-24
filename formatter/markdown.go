package formatter

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gahoolee/tmdb-cli/api"
)

func printMarkdown(data interface{}, itemType string) error {
	switch v := data.(type) {
	case *api.SearchResultPage:
		return printSearchMarkdown(v, itemType)
	case *api.MovieDetails:
		return printMovieMarkdown(v)
	case *api.TVDetails:
		return printTVMarkdown(v)
	default:
		// Fallback
		out, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf("### %s Results\n\n```json\n%s\n```\n", strings.ToUpper(itemType), string(out))
		return nil
	}
}

func printSearchMarkdown(data *api.SearchResultPage, itemType string) error {
	fmt.Printf("# Search Results (Type: %s, Total: %d, Page %d/%d)\n\n", strings.ToUpper(itemType), data.TotalResults, data.Page, data.TotalPages)
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

		fmt.Printf("%d. **%s** (%s) - Type: %s | ID: %.0f\n", idx+1, title, date, mediaType, id)
	}
	return nil
}

func printMovieMarkdown(data *api.MovieDetails) error {
	fmt.Printf("# 🎬 Movie: %s (%s)\n\n", data.Title, extractYear(data.ReleaseDate))
	if data.Tagline != "" {
		fmt.Printf("> *%s*\n\n", data.Tagline)
	}
	fmt.Printf("**Original Title:** %s\n", data.OriginalTitle)
	fmt.Printf("**Rating:** %.1f\n", data.VoteAverage)
	fmt.Printf("**Runtime:** %d minutes\n", data.Runtime)
	
	if len(data.Genres) > 0 {
		var genres []string
		for _, g := range data.Genres {
			genres = append(genres, g.Name)
		}
		fmt.Printf("**Genres:** %s\n", strings.Join(genres, ", "))
	}
	fmt.Printf("\n### Overview\n%s\n", data.Overview)
	
	if data.PosterPath != "" {
		fmt.Printf("\n### Poster\n`https://image.tmdb.org/t/p/w500%s`\n", data.PosterPath)
	}
	return nil
}

func printTVMarkdown(data *api.TVDetails) error {
	fmt.Printf("# 📺 TV: %s (%s)\n\n", data.Name, extractYear(data.FirstAirDate))
	if data.Tagline != "" {
		fmt.Printf("> *%s*\n\n", data.Tagline)
	}
	fmt.Printf("**Status:** %s\n", data.Status)
	fmt.Printf("**Rating:** %.1f\n", data.VoteAverage)
	fmt.Printf("**Seasons:** %d | **Episodes:** %d\n", data.NumberOfSeas, data.NumberOfEps)
	
	if len(data.Genres) > 0 {
		var genres []string
		for _, g := range data.Genres {
			genres = append(genres, g.Name)
		}
		fmt.Printf("**Genres:** %s\n", strings.Join(genres, ", "))
	}
	fmt.Printf("\n### Overview\n%s\n", data.Overview)
	return nil
}

func extractYear(date string) string {
	if len(date) >= 4 {
		return date[:4]
	}
	return "N/A"
}
