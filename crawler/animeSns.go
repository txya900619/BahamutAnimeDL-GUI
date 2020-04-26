package crawler

import (
	"github.com/gocolly/colly"
	"github.com/txya900619/BahamutAnimeDL-GUI/models"
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
func GetSnsByOneSn(sn string) []models.Sn {
	c := colly.NewCollector()
	Sns := make([]models.Sn, 0)
	c.OnHTML("section.season>ul>li>a", func(element *colly.HTMLElement) {
		newSn := models.Sn{}
		newSn.Sn = strings.Split(element.Attr("href"), "=")[1]
		newSn.Number = element.Text
		Sns = append(Sns, newSn)
	})
	c.Visit("https://ani.gamer.com.tw/animeVideo.php?sn=" + sn)
	return Sns
}
