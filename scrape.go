package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
)

type Menu struct {
	Title string `json:"title"`
	Price string `json:"price"`
}

func main() {

	menus := []Menu{}

	scrapeUrl := "https://www.tacobell.com/food/breakfast"

	c := colly.NewCollector(colly.AllowedDomains("www.tacobell.com", "tacobell.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("visiting %s \n", r.URL)
	})

	c.OnHTML("div.styles_card__1DpUa", func(h *colly.HTMLElement) {
		selection := h.DOM

		title := selection.Find("h4").Text()

		childNode := selection.Children().Nodes
		if len(childNode) == 4 {
			price := selection.Find("p.styles_product-details__2VdYf > span:first-child").Text()
			menu := Menu{
				Title: title,
				Price: price,
			}
			menus = append(menus, menu)
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Printf("error while scraping:%s \n", err.Error())
	})

	c.OnScraped(func(r *colly.Response) {
		jsonData, _ := json.Marshal(menus)
		fmt.Println(string(jsonData))
	})

	c.Visit(scrapeUrl)
}
