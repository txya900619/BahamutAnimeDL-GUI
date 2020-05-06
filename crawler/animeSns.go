package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/txya900619/BahamutAnimeDL-GUI/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetRealSn(ref string) string {
	resp, err := http.Get("https://ani.gamer.com.tw/animeRef.php?sn=" + ref)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(resp.Request.URL.RawQuery, "=")[1]
}
func getSnsByOneSnInOnePart(sn string) map[string][]models.Sn {
	c := colly.NewCollector()
	Sns := make([]models.Sn, 0)
	SnsOnePart := make(map[string][]models.Sn)
	c.OnHTML("section.season>ul>li>a", func(element *colly.HTMLElement) {
		newSn := models.Sn{}
		newSn.Sn = strings.Split(element.Attr("href"), "=")[1]
		newSn.Number = element.Text
		Sns = append(Sns, newSn)
	})
	c.Visit("https://ani.gamer.com.tw/animeVideo.php?sn=" + sn)
	if len(Sns) == 0 {
		Sns = append(Sns, models.Sn{Sn: sn, Number: "1"})
	}
	SnsOnePart[""] = Sns
	return SnsOnePart
}

func getSnsByOneSnInTwoPart(sn string) map[string][]models.Sn {
	c := colly.NewCollector()
	SnsTwoPart := make(map[string][]models.Sn)
	c.OnHTML("section.season>p", func(element *colly.HTMLElement) {
		Sns := make([]models.Sn, 0)
		element.DOM.Next().Children().Each(func(i int, selection *goquery.Selection) {
			a := selection.Children()
			newSn := models.Sn{}
			href, _ := a.Attr("href")
			newSn.Sn = strings.Split(href, "=")[1]
			newSn.Number = a.Text()
			Sns = append(Sns, newSn)
		})
		SnsTwoPart[element.Text] = Sns
	})
	c.Visit("https://ani.gamer.com.tw/animeVideo.php?sn=" + sn)
	return SnsTwoPart
}

func GetSnsByOneSn(sn string) map[string][]models.Sn {
	resp, err := http.Get("https://ani.gamer.com.tw/animeVideo.php?sn=" + sn)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string][]models.Sn)
	if strings.Contains(string(body), "本篇") {
		result = getSnsByOneSnInTwoPart(sn)
	} else {
		result = getSnsByOneSnInOnePart(sn)
	}
	return result
}
