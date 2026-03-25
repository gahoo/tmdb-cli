# TMDB CLI

[English](README.md) | [简体中文](README_ZH.md)


A powerful command-line interface tool for querying [The Movie Database (TMDB)](https://www.themoviedb.org/) API. Built with Go, it allows you to search for movies, TV shows, and people, as well as export data in various formats like JSON, Markdown, and NFO (Kodi/Jellyfin compatible).

## Features

- **Search**: Find movies, TV shows, and people with multi-search support.
- **Detailed Info**: Retrieve comprehensive data for movies, TV series, seasons, and individual episodes.
- **Trending**: Check what's currently trending globally.
- **Collections**: Get details about movie collections.
- **Find by External ID**: Look up items using IMDb, TVDB, or other external IDs.
- **Multiple Output Formats**:
  - `json`: Standard JSON output (supports field filtering).
  - `markdown`: Beautifully formatted Markdown.
  - `nfo`: Kodi/Jellyfin compatible XML metadata files.
  - `table`: Clean CLI table view.
- **Poster Download**: Automatically download movie/series posters when exporting to NFO.
- **Persistent Configuration**: Save your API token and preferred language.

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/gahoolee/tmdb-cli.git
cd tmdb-cli
go build -o tmdb
```

## Configuration

Before using the tool, you need to set your TMDB API token (v4 Read Access Token or v3 API Key).

```bash
# Set your API token
./tmdb config set-auth YOUR_TMDB_API_TOKEN

# Set default language (optional, e.g., zh-CN, en-US)
./tmdb config set-lang zh-CN
```

The configuration is saved in `~/.tmdb.json`.

## Usage

### General Flags

- `--format`: Output format (json, markdown, nfo, table). Default: `json`.
- `--output`, `-o`: Save output to a specific file.
- `--fields`, `-f`: Comma-separated list of fields to include (e.g., `title,overview,budget`).
- `--language`, `-l`: Override default language for this request.
- `--poster`: Download poster image locally (only applicable for `nfo` format).

### Examples

#### Search for a Movie
```bash
./tmdb search "Inception" --type movie --format table
```

#### Get Movie Details and Download NFO + Poster
```bash
./tmdb movie 27205 --format nfo --poster
```

#### Get TV Series Details
```bash
./tmdb tv 60625 --format markdown
```

#### Get a Specific TV Season or Episode
```bash
# Get Season 1 of a TV show
./tmdb tv 60625 --season 1

# Get Episode 1 of Season 1
./tmdb tv 60625 --season 1 --episode 1 --format markdown
```

#### Find by External ID (IMDb)
```bash
./tmdb find tt0133093 --source imdb_id --format table
```

#### View Trending Movies
```bash
./tmdb trending --type movie --time day --format table
```

#### Get Collection Details
```bash
./tmdb collection 10 --format markdown
```

## Output Formats

### NFO Generation
When using `--format nfo`, the tool generates XML metadata files compatible with media managers like Kodi, Jellyfin, and Emby. 
- Movies generate a `.nfo` file.
- TV Shows generate `tvshow.nfo`.
- Seasons generate `season.nfo`.
- Episodes generate a `.nfo` file.

Combined with the `--poster` flag, it helps in setting up local media libraries quickly.

## Requirements
- Go 1.26 or higher.

## License
MIT License
