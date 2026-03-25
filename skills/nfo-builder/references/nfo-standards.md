# Kodi / Jellyfin NFO Standards

## Core Rules
1. **Format**: XML text file with `.nfo` extension.
2. **Encoding**: Must be saved with `UTF-8` encoding.
3. **Empty Tags**: If data is missing for an optional tag, omit the tag rather than leaving it empty, unless using it as a template placeholder.
4. **Dates**: Always use the format `yyyy-mm-dd` for `<premiered>` and `<aired>`. Avoid using the deprecated `<year>` tag.

## Naming Conventions & File Locations
| Media Type | Recommended NFO Filename | Location |
| :--- | :--- | :--- |
| **Movie** | `<VideoFileName>.nfo` | Next to the video file. |
| **TV Show** | `tvshow.nfo` | At the root of the TV show folder. |
| **TV Season** | `season.nfo` | At the root of the Season folder (e.g., `Season 1/`). |
| **Episode** | `<VideoFileName>.nfo` | Next to the video file. |
| **Music Artist**| `artist.nfo` | Root of the Artist folder. |
| **Music Album** | `album.nfo` | Root of the Album folder. |

## Important Tags

### The `<uniqueid>` Tag
This is critical for scraper matching.
- **Attributes**: `type` (e.g., `"imdb"`, `"tmdb"`, `"tvdb"`) and `default` (`"true"` or `"false"`).
- **Multiple Allowed**: You can include multiple `<uniqueid>` tags, but only ONE can have `default="true"`.
- **Example**:
  ```xml
  <uniqueid type="tmdb" default="true">12345</uniqueid>
  <uniqueid type="imdb" default="false">tt1234567</uniqueid>
  ```

### TV Season Specific Tags (`season.nfo`)
- `<showtitle>`: The name of the TV Series.
- `<seasonnumber>`: The numerical index of the season.
- `<plot>`: A summary of the season (can be concatenated from episode titles if needed).

### The `<episodeguide>` Tag (TV Shows)
In Kodi v19 and later, this tag uses JSON to map TV show IDs for episode scraping.
- **Example**:
  ```xml
  <episodeguide>
    {"tmdb": "76479", "imdb": "tt1190634", "tvdb": "355567"}
  </episodeguide>
  ```

### Localized Metadata
The `tmdb-cli` supports fetching localized metadata via the `--language` flag.
- **Rule**: Ensure the NFO file is saved in `UTF-8` to support non-ASCII characters (e.g., Chinese, Japanese).
- **Example**: `./tmdb movie 27205 --language zh-CN --format nfo`

### Combination NFO
If a user just wants to scrape the data from a site but override *some* specific tags (like `<genre>`), the NFO should contain the XML data to override, followed by the URL of the scraper page at the very end of the file.
- **Example**:
  ```xml
  <movie>
      <title>Custom Title</title>
      <genre clear="true">Sci-Fi</genre>
  </movie>
  https://www.themoviedb.org/movie/11-star-wars
  ```
  *(The `clear="true"` attribute tells Kodi to clear out the scraper's tags and only use the local ones for that specific element).*