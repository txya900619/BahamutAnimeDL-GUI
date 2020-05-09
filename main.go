// +build !dev

package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/markbates/pkger"
	"github.com/txya900619/BahamutAnimeDL-GUI/crawler"
	"github.com/txya900619/BahamutAnimeDL-GUI/database"
	dbModels "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/queue"
	"github.com/txya900619/BahamutAnimeDL-GUI/utilities"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/zserge/lorca"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var NewAnimeList []models.NewAnime
var AnimeList []models.Anime
var db *sql.DB
var queueSystem *queue.System

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

	queueSystem = queue.New(db, 1)
	go queueSystem.Start()
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

	app.Bind("getNewAnimeList", func() []models.NewAnime {
		return NewAnimeList
	})

	app.Bind("getAnimesByPage", func(page int) []models.Anime {
		return AnimeList[(page-1)*18 : page*18]
	})
	app.Bind("getMaxPage", func() int {
		return len(AnimeList)/18 + 1
	})

	app.Bind("getAnimesByFilter", func(filter string) []models.Anime {
		filteredAnimes := make([]models.Anime, 0)
		filter = strings.ToLower(filter)
		for _, v := range AnimeList {
			if strings.Contains(strings.ToLower(v.Title), filter) {
				filteredAnimes = append(filteredAnimes, v)
			}
		}
		return filteredAnimes
	})

	app.Bind("getRealSn", func(ref string) string {
		return crawler.GetRealSn(ref)
	})

	app.Bind("getAnimeAllSn", func(sn string) map[string][]models.Sn {
		return crawler.GetSnsByOneSn(sn)
	})

	app.Bind("insertAnimeToQueue", func(title, ep, sn string, spacial bool) {
		lastSequence, err := dbModels.DownloadQueues().Count(context.Background(), db)
		if err != nil {
			log.Fatal(err)
		}
		var intSpacial int64
		if spacial {
			intSpacial = 1
		}
		intSn, err := strconv.ParseInt(sn, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		queue := dbModels.DownloadQueue{SN: intSn, Name: title, Ep: ep, Sequence: lastSequence + 1, Spacial: intSpacial}
		err = queue.Insert(context.Background(), db, boil.Infer())
	})

	net, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer net.Close()

	go http.Serve(net, http.FileServer(pkger.Dir("/dist")))
	app.Load(fmt.Sprintf("http://%s", net.Addr()))
	<-app.Done()
}
