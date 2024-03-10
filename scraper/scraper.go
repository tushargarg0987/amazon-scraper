package scraper

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly/v2"
)

type product struct {
	Brand    string `json:"brand"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	ImageUrl string `json:"imageurl"`
	Rating   string `json:"rating"`
}

func Scrap(search string) ([]byte, error) {
	ProductList := []product{}
	c := colly.NewCollector()

	c.OnHTML("div.puis-card-container", func(e *colly.HTMLElement) {
		brand := e.ChildText("div[data-cy=title-recipe] > div.a-row > h2.a-size-mini > span.a-size-base-plus")
		name := e.ChildText("h2.s-line-clamp-2 > a.s-underline-link-text > span.a-text-normal")
		if name != "" && brand != "" {
			newProd := new(product)
			newProd.Name = e.ChildText("h2.s-line-clamp-2 > a.s-underline-link-text > span.a-text-normal")
			newProd.Brand = e.ChildText("div[data-cy=title-recipe] > div.a-row > h2.a-size-mini > span.a-size-base-plus")
			newProd.Price = e.ChildText("span.a-price > span > span.a-price-whole")
			newProd.Rating = e.ChildAttr("div.a-size-small > span[aria-label]", "aria-label")
			newProd.ImageUrl = e.ChildAttr("img.s-image", "src")
			ProductList = append(ProductList, *newProd)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.amazon.in/s?k=" + search)
	jsonData, _ := json.Marshal(ProductList)
	return jsonData, nil
}
