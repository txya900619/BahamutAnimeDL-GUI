package crawler

import "github.com/gocolly/colly"

func Check18Up(sn string) bool {
	is18Up := false
	c := colly.NewCollector()
	c.OnHTML("div.rating>img", func(element *colly.HTMLElement) {
		if element.Attr("src") == "https://i2.bahamut.com.tw/acg/TW-18UP.gif" {
			is18Up = true
		}
	})
	c.Visit("https://ani.gamer.com.tw/animeVideo.php?sn=" + sn)
	return is18Up
}
