---
name: tmdb-nfo-builder
description: Generate Kodi and Jellyfin compatible .nfo XML metadata files mainly as a fallback for the tmdb-cli project.
---
# NFO Generator Skill

This skill is designed to work in conjunction with the `tmdb-cli` project to generate Kodi and Jellyfin compatible `.nfo` metadata files.

## Workflow 1: Try TMDB API first (Primary)

Before using manual templates, attempt to fetch information and generate the NFO file directly using the `tmdb-cli` tool.

1.  **Search for the Media ID**: Use `markdown` format to clearly see the search results and IDs.
    ```bash
    ./tmdb search "Inception" --type movie --format markdown
    ```
2.  **Generate NFO**: Use the ID to generate the NFO file. The `--output` flag **must** specify the full filename, not just a path. Use `--poster` to download the artwork.
    - **Movie**: `./tmdb movie <ID> --format nfo --poster --output "/path/to/movie/MovieName.nfo"`
    - **TV Show**: `./tmdb tv <ID> --format nfo --poster --output "/path/to/show/tvshow.nfo"`
    - **TV Season**: `./tmdb tv <ID> --season <number> --format nfo --output "/path/to/show/Season X/season.nfo"`
    - **TV Episode**: `./tmdb tv <ID> --season <number> --episode <number> --format nfo --output "/path/to/show/Season X/EpisodeName.nfo"`
3.  **Examples**:
    - Search and generate for a movie:
      `./tmdb search "The Matrix" --type movie --format markdown` (Find ID 603)
      `./tmdb movie 603 --format nfo --poster --output "/Users/user/Movies/The Matrix (1999)/The Matrix (1999).nfo"`
    - Generate for a specific TV Season:
      `./tmdb tv 60625 --season 1 --format nfo --output "/Users/user/TV/The Flash (2014)/Season 1/season.nfo"`

If `tmdb-cli` successfully retrieves the information and writes the required `.nfo` file, you are done.

## Workflow 2: Fallback to Manual Generation (When TMDB Fails)

Only when `tmdb-cli` **fails** to find the media or returns no usable information, follow this manual fallback process:

1.  **Identify Media Type**: Determine if it's a Movie, TV Show, Season, Episode, etc.
2.  **Search for Metadata**: Perform a web search (IMDb, Wikipedia, Douban, etc.) to gather: Title, Year, Plot, Genres, Ratings, Posters.
3.  **Load Template**: Read the appropriate XML template from `assets/templates/`.
4.  **Populate Data**: Fill in the XML tags based on the gathered metadata.
5.  **User Review (Mandatory)**: Present the generated XML content to the user for review.
6.  **Output**: **Only after user approval**, write the generated XML using the correct NFO naming convention (see `references/nfo-standards.md`).

## Workflow 3: Directory-Based Generation (Final Fallback)

If no information is found via TMDB or Web Search:

1.  **Use Directory/File Names**: Extract the title and year from the parent directory name or the filename itself.
    - Example: `/Movies/Inception (2010)/` -> Title: "Inception", Year: "2010".
2.  **Season Plot Concatenation**: For TV Seasons, if individual episode information is available through filenames, generate a summary plot for the season by concatenating the episode names or summaries.
3.  **Simple NFO**: Create a minimal NFO using the `movie.xml`, `tvshow.xml`, or `season.xml` template with the `<title>` and `<uniqueid type="local">` tags.
4.  **User Review (Mandatory)**: Present the generated XML content to the user for review.
5.  **Output**: **Only after user approval**, write the generated XML using the correct NFO naming convention.

## Special Case: Combination NFOs
(See `references/nfo-standards.md` for more details).
Format:
```xml
<movie>
  <title>Custom Title</title>
  <genre clear="true">Custom Genre</genre> 
</movie>
https://www.themoviedb.org/movie/12345
```