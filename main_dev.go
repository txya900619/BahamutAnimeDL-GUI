// +build dev

package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/txya900619/BahamutAnimeDL-GUI/crawler"
	"github.com/txya900619/BahamutAnimeDL-GUI/database"
	dbModels "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/utilities"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/zserge/lorca"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var NewAnimeList []models.NewAnime
var AnimeList []models.Anime
var db *sql.DB

func init() {
	db = database.ConnectSqlite()
	NewAnimeList = crawler.GetNewAnimeList()
	AnimeList = crawler.GetAllAnimeList()

	if _, err := os.Stat("./.temp"); os.IsNotExist(err) {
		os.Mkdir("./.temp", 0777)
		if runtime.GOOS == "windows" {
			utilities.HideFolder("./.temp")
		}
	}

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

	app.Bind("insertAnimeToQueue", func(title, ep, sn string) {
		lastSequence, err := dbModels.DownloadQueues().Count(context.Background(), db)
		if err != nil {
			log.Fatal(err)
		}
		intSn, err := strconv.ParseInt(sn, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		queue := dbModels.DownloadQueue{SN: intSn, Name: title, Ep: ep, Sequence: lastSequence + 1}
		err = queue.Insert(context.Background(), db, boil.Infer())
	})

	app.Load(fmt.Sprintf("http://%s", "127.0.0.1:8080"))
	<-app.Done()
}
