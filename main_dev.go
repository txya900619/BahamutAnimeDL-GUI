//+build dev

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/zserge/lorca"
	"log"
	"runtime"
	"strings"
)

type Anime struct {
	Title string   `json:"title"`
	Img   string   `json:"img"`
	Sn    []string `json:"sn"`
	Vol   int      `json:"vol"`
}

func main() {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	app, err := lorca.New("", "", 1200, 800, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()
	app.Bind("getAnimeList", getAnimeList)

	app.Load(fmt.Sprintf("http://%s", "127.0.0.1:8080"))

	<-app.Done()
}
func getAnimeList() string {
	animeList := make([]Anime, 0)
	c := colly.NewCollector(
		colly.Async(true),
	)
	playPageCollector := c.Clone()
	c.OnHTML("div.index_season.view-grid[style!='display:none']", func(element *colly.HTMLElement) {
		element.ForEach(".newanime", func(index int, element *colly.HTMLElement) {
			content := element.DOM.Children().Eq(1)
			title := content.Children().Eq(1).Children().Eq(0).Text()
			for _, v := range animeList {
				if v.Title == title {
					return
				}
			}
			sn, _ := content.Attr("href")
			img, _ := content.Children().Children().Attr("data-src")
			anime := Anime{
				Title: title,
				Img:   img,
			}
			anime.Sn = append(anime.Sn, strings.Split(sn, "=")[1])
			animeList = append(animeList, anime)
			element.Request.Ctx.Put(sn, len(animeList)-1)
			playPageCollector.Request("GET", sn, nil, element.Request.Ctx, nil)
		})
		playPageCollector.Wait()

	})
	playPageCollector.OnHTML(".season", func(element *colly.HTMLElement) {
		sn := element.ChildAttrs("a[href]", "href")
		for i, childSn := range sn {
			sn[i] = strings.Split(childSn, "=")[1]
		}
		index := element.Request.Ctx.GetAny(element.Request.URL.String()).(int)
		animeList[index].Sn = sn
	})
	c.Visit("https://ani.gamer.com.tw/")
	c.Wait()

	for i, v := range animeList {
		animeList[i].Vol = len(v.Sn)
	}
	jsonAnimeList, _ := json.Marshal(animeList)
	return string(jsonAnimeList)
}
