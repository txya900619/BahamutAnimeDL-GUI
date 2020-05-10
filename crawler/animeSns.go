package crawler

import (
	"context"
	"database/sql"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	dbModels "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetRealSn(ref string) string {
	resp, err := http.Get("https://ani.gamer.com.tw/animeRef.php?sn=" + ref)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(resp.Request.URL.RawQuery, "=")[1]
}
func getSnsByOneSnInOnePart(title, sn string, db *sql.DB) map[string][]models.Sn {
	c := colly.NewCollector()
	Sns := make([]models.Sn, 0)
	SnsOnePart := make(map[string][]models.Sn)
	c.OnHTML("section.season>ul>li>a", func(element *colly.HTMLElement) {
		newSn := models.Sn{CanDownload: true}
		newSn.Sn = strings.Split(element.Attr("href"), "=")[1]
		newSn.Number = element.Text

		intSn, _ := strconv.ParseInt(newSn.Sn, 10, 64)
		if _, err := dbModels.FindDownloadQueue(context.Background(), db, intSn); err == nil {
			newSn.CanDownload = false
		}
		if _, err := os.Stat("./download/" + title + "/" + title + " [" + newSn.Number + "].mp4"); err == nil {
			newSn.CanDownload = false
		}
		Sns = append(Sns, newSn)
	})
	c.Visit("https://ani.gamer.com.tw/animeVideo.php?sn=" + sn)
	if len(Sns) == 0 {
		Sns = append(Sns, models.Sn{Sn: sn, Number: "1"})
	}
	SnsOnePart[""] = Sns
	return SnsOnePart
}

func getSnsByOneSnInTwoPart(title, sn string, db *sql.DB) map[string][]models.Sn {
	c := colly.NewCollector()
	SnsTwoPart := make(map[string][]models.Sn)
	c.OnHTML("section.season>p", func(element *colly.HTMLElement) {
		Sns := make([]models.Sn, 0)
		element.DOM.Next().Children().Each(func(i int, selection *goquery.Selection) {
			a := selection.Children()
			newSn := models.Sn{CanDownload: true}
			href, _ := a.Attr("href")
			newSn.Sn = strings.Split(href, "=")[1]
			newSn.Number = a.Text()

			intSn, _ := strconv.ParseInt(newSn.Sn, 10, 64)
			if _, err := dbModels.FindDownloadQueue(context.Background(), db, intSn); err == nil {
				newSn.CanDownload = false
			}
			if element.Text == "特別篇" {
				if _, err := os.Stat("./download/" + title + "/" + title + " 特別篇 [" + newSn.Number + "].mp4"); err == nil {
					newSn.CanDownload = false
				}
			} else {
				if _, err := os.Stat("./download/" + title + "/" + title + " [" + newSn.Number + "].mp4"); err == nil {
					newSn.CanDownload = false
				}
			}

			Sns = append(Sns, newSn)
		})
		SnsTwoPart[element.Text] = Sns
	})
	c.Visit("https://ani.gamer.com.tw/animeVideo.php?sn=" + sn)
	return SnsTwoPart
}

func GetSnsByOneSn(title, sn string, db *sql.DB) map[string][]models.Sn {
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
		result = getSnsByOneSnInTwoPart(title, sn, db)
	} else {
		result = getSnsByOneSnInOnePart(title, sn, db)
	}
	return result
}
