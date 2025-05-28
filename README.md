# webcrawler
Go-based crawler and scraper using the Colly framework

## Using webcrawler
The application can be run from the terminal.  Clone the repository, then navigate to the folder containing the repository.  The executable can be run with
```
webcrawler
```
Alternatively, the webcrawler.go file can be run directly with
```
go run webcrawler.go
```

## About
webcrawler crawls and scrapes ten Wikipedia pages for text content.  The pages' URLs can be found (and adjusted if desired) in the `main` function in `webcrawler.go`:
```
urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}
```

### Output
webcrawler produces a JSON lines file `wikipedia_entries.jsonl` with one line per URL scraped.  The JSON tags are goverened by the type `wikientry` in `webcrawler.go`:
```
type wikiEntry struct {
	URL   string   `json:"url"`
	Title string   `json:"title"`
	Links []string `json:"links"`
	Body  string   `json:"body"`
}
```

### Known Issues
Some HTML tags containing style or metadata content are being included in the `"body"` value.  This has been attempted to be addressed by skipping certain HTML tags like `<style>`:
```
c.OnHTML("div.mw-body-content", func(e *colly.HTMLElement) {
		e.ForEach("*", func(_ int, child *colly.HTMLElement) {
			// omit skipped tags
			switch child.Name {
			case "style":
				return
			default:
				entry.Body = entry.Body + "\n\n" + child.Text
			}

		})
  ...
```
This is not yet having the desired effect; `<style>` tags are still being added to the struct `entry` in the field `entry.Body`
