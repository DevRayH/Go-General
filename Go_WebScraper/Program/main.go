package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func stringFormater(rawString string) []string { 
	var itemSlice[]string = strings.Split(rawString, "#")

	return itemSlice
}


func main() {
	// Entry point
	c := colly.NewCollector()

	//div[class=a-carousel-row-inner]
	c.OnHTML("div[id=zg_left_col1]", func(h *colly.HTMLElement){
		// Process the HTML element
		//fmt.Println("Item found\n\n")
		
		var stringSlice[]string = stringFormater(h.ChildText("div[class=a-carousel-row-inner]"))
		
		for _, item := range stringSlice {
			fmt.Printf("%v \n\n", item)
		}
	})

	c.Visit(stringMap["siteURL"])
}