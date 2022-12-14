package web_scrap

import (
	"errors"
	"github.com/gocolly/colly"
)

type Scrape struct {
}

func (Scrape) Demo() (value []interface{}, err error) {
	c := colly.NewCollector()
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			value = append(value,
				[]string{
					el.ChildText("td:nth-child(1)"),
					el.ChildText("td:nth-child(2)"),
					el.ChildText("td:nth-child(3)"),
				})
		})
	})
	err = c.Visit("https://www.w3schools.com/html/html_tables.asp")
	if err != nil {
		return nil, errors.New("Error: " + err.Error())
	}
	return value, nil
}
