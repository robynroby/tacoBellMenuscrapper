package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Menu struct{
	Title string
	// MenuItem string
	Price float64
}

func main() {

	// menu := Menu{}

	scrapeUrl := "https://www.tacobell.com/food/breakfast"

	c := colly.NewCollector(colly.AllowedDomains("www.tacobell.com", "tacobell.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s \n", r.URL)
	})

	c.OnHTML("div.styles_card__1DpUa", func(h *colly.HTMLElement) {
		selection := h.DOM

		title := selection.Find("h4").Text()
		// price := selection.Find("p.styles_product-details__2VdYf > span:first-child")
		fmt.Printf("%s : ",title)

		childNode := selection.Children().Nodes
		if len(childNode) == 4 {
			price := selection.Find("p.styles_product-details__2VdYf > span:first-child").Text()
			
			fmt.Printf("%s \n",price)
		} 

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("error while scraping:%s \n", err.Error())
	})

	c.Visit(scrapeUrl)
}
