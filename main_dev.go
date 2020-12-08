// +build dev

package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/webview/webview"

	"github.com/txya900619/BahamutAnimeDL-GUI/database"
	dbModels "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/queue"
	"github.com/txya900619/BahamutAnimeDL-GUI/utilities"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var db *sql.DB
var queueSystem *queue.System
var maxDownloader int

func init() {
	db = database.ConnectSqlite()
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

	app := webview.New(true)
	defer app.Destroy()
	app.SetTitle("")
	app.SetSize(1200, 800, webview.HintNone)

	app.Bind("insertAnimeToQueue", func(title, ep, sn string, spacial bool) bool {

		//TODO: 18
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

	app.Navigate(fmt.Sprintf("http://%s", "127.0.0.1:8080"))
}
