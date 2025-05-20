package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"encoding/json"

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
			fmt.Println("URL not found", url)
		}
		err = write(entry, file)
		if err != nil {
			fmt.Println("Problem converting to JSON and/or writing to file")
		}
	}

	// temp print to use c
	fmt.Println(c)
}

type wikiEntry struct {
	URL   string   `json:"url"`
	Title string   `json:"title"`
	Links []string `json:"links"`
	Body  string   `json:"body"`
}

func write(w wikiEntry, f *os.File) error {
	b, err := json.Marshal(w)
	if err != nil {
		fmt.Println("Error writing JSON to file")
		return err
	}
	fmt.Fprintln(f, b)
	return nil
}

func urlToWikiEntry(c *colly.Collector, url string) (wikiEntry, error) {
	var entry wikiEntry
	entry.URL = url
	// Wikipedia urls of the form https://en.wikipedia.org/wiki/Title:
	entry.Title = strings.Split(url, "/")[4]
	c.OnHTML("div.mw-body-content", func(e *colly.HTMLElement) {
		entry.Body = e.Text
		e.ForEach("a[href]", func(_ int, link *colly.HTMLElement) {
			entry.Links = append(entry.Links, link.Attr("href"))
		})

	})
	err := c.Visit(url)
	return entry, err
}
