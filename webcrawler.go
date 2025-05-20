package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	fName := "wikipedia_entries.jsonl"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"))
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
	for _, url := range urls {
		entry, err := urlToWikiEntry(c, url)
		if err != nil {
			// error handling here
		}
		// temp print to console to use entry
		fmt.Println(entry)
	}

	// temp print to use c
	fmt.Println(c)
}

type wikiEntry struct {
	url   string
	title string
	links []string
	body  string
}

func urlToWikiEntry(c *colly.Collector, url string) (wikiEntry, error) {
	var entry wikiEntry
	entry.url = url
	// Wikipedia urls of the form https://en.wikipedia.org/wiki/Title:
	entry.title = strings.Split(url, "/")[4]
	c.OnHTML("div.mw-body-content", func(e *colly.HTMLElement) {
		entry.body = e.Text
		e.ForEach("a[href]", func(_ int, link *colly.HTMLElement) {
			entry.links = append(entry.links, link.Attr("href"))
		})

	})
	err := c.Visit(url)
	return entry, err
}

func wikiEntryToJsonl(e wikiEntry) (string, error) {
	// temp return
	fmt.Println(e)
	return "tacos", errors.New("temp error")
}
