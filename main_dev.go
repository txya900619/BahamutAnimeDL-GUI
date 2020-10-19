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
	"github.com/txya900619/BahamutAnimeDL-GUI/queue"
	"github.com/txya900619/BahamutAnimeDL-GUI/utilities"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zserge/lorca"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var NewAnimeList []models.NewAnime
var AnimeList []models.Anime
var db *sql.DB
var queueSystem *queue.System
var maxDownloader int

func init() {
	db = database.ConnectSqlite()
	NewAnimeList = crawler.GetNewAnimeList()
	AnimeList = crawler.GetAllAnimeList()
	maxDownloader = 1

	if _, err := os.Stat("./.temp"); os.IsNotExist(err) {
		os.Mkdir("./.temp", 0777)
		if runtime.GOOS == "windows" {
			utilities.HideFolder("./.temp")
		}
	}

	queueSystem = queue.New(db, maxDownloader)
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
		filter = strings.ToLower(filter)
		filteredAnimes := make([]models.Anime, 0)
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

	app.Bind("getAnimeAllSn", func(title, sn string) map[string][]models.Sn {
		return crawler.GetSnsByOneSn(title, sn, db)
	})

	app.Bind("insertAnimeToQueue", func(title, ep, sn string, spacial bool) bool {
		if crawler.Check18Up(sn) {
			return false
		}
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
		return true
	})

	app.Bind("getDownloadQueue", func() models.QueueStatus {
		queuesPtr, err := dbModels.DownloadQueues(qm.OrderBy("sequence")).All(context.Background(), db)
		if err != nil {
			log.Fatal(err)
		}
		queues := make([]dbModels.DownloadQueue, 0)
		downloadStatus := make([]int, 0)
		for _, queuePtr := range queuesPtr {
			if queuePtr.Downloading == 1 {
				files, _ := ioutil.ReadDir("./.temp/" + strconv.FormatInt(queuePtr.SN, 10))
				downloadStatus = append(downloadStatus, len(files))
			}
			queues = append(queues, *queuePtr)
		}
		return models.QueueStatus{queues, downloadStatus}
	})

	app.Load(fmt.Sprintf("http://%s", "127.0.0.1:8080"))
	<-app.Done()
}
