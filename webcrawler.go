package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	fName := "xkcd_store_items.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	
	c := colly.NewCollector()
	urls := "temp"
	fmt.Println(c)
}
