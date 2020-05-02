//+build dev

package main

import (
	"fmt"
	"github.com/txya900619/BahamutAnimeDL-GUI/crawler"
	"github.com/txya900619/BahamutAnimeDL-GUI/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/utilities"
	"github.com/zserge/lorca"
	"log"
	"runtime"
	"strings"
)

var NewAnimeList []models.NewAnime
var AnimeList []models.Anime

func init() {
	NewAnimeList = crawler.GetNewAnimeList()
	AnimeList = crawler.GetAllAnimeList()

}

func main() {

	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	args = append(args, "--disable-features=TranslateUI")
	app, err := lorca.New("", "", 1200, 800, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	app.Bind("getNewAnimeList", func() string {
		return utilities.ToJson(NewAnimeList)
	})
	app.Bind("getAllAnimeList", func() string {
		return utilities.ToJson(AnimeList)
	})
	app.Bind("getAnimesByPage", func(page int) string {
		return utilities.ToJson(AnimeList[(page-1)*18 : page*18])
	})
	app.Bind("getMaxPage", func() int {
		return len(AnimeList)/18 + 1
	})

	app.Bind("getAnimesByFilter", func(filter string) string {
		filter = strings.ToLower(filter)
		filteredAnimes := make([]models.Anime, 0)
		for _, v := range AnimeList {
			if strings.Contains(strings.ToLower(v.Title), filter) {
				filteredAnimes = append(filteredAnimes, v)
			}
		}
		return utilities.ToJson(filteredAnimes)
	})

	app.Bind("getRealSn", func(ref string) string {
		return crawler.GetRealSn(ref)
	})

	app.Bind("getAnimeAllSn", func(sn string) string {
		return utilities.ToJson(crawler.GetSnsByOneSn(sn))
	})

	app.Load(fmt.Sprintf("http://%s", "127.0.0.1:8080"))
	<-app.Done()
}
