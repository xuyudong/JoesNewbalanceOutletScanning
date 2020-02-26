package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

func main()  {
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.joesnewbalanceoutlet.com"),
	)

/*	c.OnHTML("#Items img[src^='https://s7']", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})*/
	c.OnHTML("#Items .figureWrapper", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		fmt.Printf("Link found: id=%s name=%q -> %s ->%s\n",
			e.ChildAttr("a","data-productcode"),
			e.ChildAttr("a","data-productname"),
			e.ChildAttr("img[src^='https://s7']","src"),
			e.ChildAttr("a","data-price")+"$",
		)
		//fmt.Println(e.ChildAttr("a","data-productname"))
		c.Visit(e.Request.AbsoluteURL(link))
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("爬虫url", r.URL.String())
	})
	for i := 1; i < 30; i++ {
		//c.Visit("https://www.joesnewbalanceoutlet.com/men/featured?Page="+strconv.Itoa(i))
		c.Visit("https://www.joesnewbalanceoutlet.com/products/?Text=997&Page="+strconv.Itoa(i))
	}
}