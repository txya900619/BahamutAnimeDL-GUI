// +build !dev

package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/markbates/pkger"
	"github.com/txya900619/BahamutAnimeDL-GUI/database"
	dbModels "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/queue"
	"github.com/txya900619/BahamutAnimeDL-GUI/utilities"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/webview/webview"
)

var db *sql.DB
var queueSystem *queue.System

func init() {
	db = database.ConnectSqlite()

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
	app := webview.New(true)
	defer app.Destroy()
	app.SetSize(1200, 800, webview.HintNone)
	app.SetTitle("")

	app.Bind("insertAnimeToQueue", func(title, ep, sn string, spacial bool) bool {
		//TODO: maybe test 18

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
		downloadQueue := dbModels.DownloadQueue{SN: intSn, Name: title, Ep: ep, Sequence: lastSequence + 1, Spacial: intSpacial}
		err = downloadQueue.Insert(context.Background(), db, boil.Infer())
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
		return models.QueueStatus{Queue: queues, DownloadingStatus: downloadStatus}
	})

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	go http.Serve(listener, http.FileServer(pkger.Dir("/dist")))
	app.Navigate(fmt.Sprintf("http://%s", listener.Addr()))

}
