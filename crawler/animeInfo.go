package crawler

import (
	"github.com/gocolly/colly"
	"github.com/txya900619/BahamutAnimeDL-GUI/model"
	"strings"
)

func GetAllAnimeList() []model.Anime {
	animeList := make([]model.Anime, 0)
	c := colly.NewCollector(colly.Async(true))
	c.OnHTML("ul.anime_list>li", func(element *colly.HTMLElement) {
		anime := model.Anime{}
		anime.Ref = strings.Split(element.ChildAttr("a", "href"), "=")[1]
		anime.Img = element.ChildAttr("a>div.pic", "data-bg")
		anime.Title = element.ChildText("div.info>b")
		animeList = append(animeList, anime)
	})
	c.OnHTML("a.next", func(element *colly.HTMLElement) {
		if element.Attr("href") == "javascript:alert('沒有下一頁');" {
			return
		}
		c.Visit("https://ani.gamer.com.tw/animeList.php" + element.Attr("href"))
	})

	c.Visit("https://ani.gamer.com.tw/animeList.php")
	c.Wait()
	return animeList
}
func GetNewAnimeList() []model.NewAnime {
	animeList := make([]model.NewAnime, 0)
	c := colly.NewCollector(
		colly.Async(true),
	)
	c.OnHTML("div.index_season.view-grid[style!='display:none']", func(element *colly.HTMLElement) {
		element.ForEach(".newanime", func(index int, element *colly.HTMLElement) {
			anime := model.NewAnime{}
			content := element.DOM.Children().Eq(1)
			anime.Title = content.Children().Eq(1).Children().Eq(0).Text()
			for _, v := range animeList {
				if v.Title == anime.Title {
					return
				}
			}
			anime.Img, _ = content.Children().Children().Attr("data-src")
			sn, _ := content.Attr("href")
			anime.Sn = strings.Split(sn, "=")[1]
			animeList = append(animeList, anime)
		})

	})
	c.Visit("https://ani.gamer.com.tw/")
	c.Wait()
	return animeList
}
