// +build !dev

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/markbates/pkger"
	"github.com/txya900619/BahamutAnimeDL-GUI/queue"
	"github.com/webview/webview"
)

var db *sql.DB
var queueSystem *queue.System

// func init() {
// 	db = database.ConnectSqlite()

// 	if _, err := os.Stat("./.temp"); os.IsNotExist(err) {
// 		os.Mkdir("./.temp", 0777)
// 		if runtime.GOOS == "windows" {
// 			utilities.HideFolder("./.temp")
// 		}
// 	}
// 	queueSystem = queue.New(db, 1)
// 	go queueSystem.Start()

// }

func main() {
	app := webview.New(true)
	defer app.Destroy()
	app.SetSize(1200, 800, webview.HintNone)
	app.SetTitle("")
	// app.Bind("insertAnimeToQueue", func(title, ep, sn string, spacial bool) bool {
	// 	//TODO: maybe test 18

	// 	lastSequence, err := dbModels.DownloadQueues().Count(context.Background(), db)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	var intSpacial int64
	// 	if spacial {
	// 		intSpacial = 1
	// 	}
	// 	intSn, err := strconv.ParseInt(sn, 10, 64)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	downloadQueue := dbModels.DownloadQueue{SN: intSn, Name: title, Ep: ep, Sequence: lastSequence + 1, Spacial: intSpacial}
	// 	err = downloadQueue.Insert(context.Background(), db, boil.Infer())
	// 	return true
	// })

	// app.Bind("getDownloadQueue", func() models.QueueStatus {
	// 	queuesPtr, err := dbModels.DownloadQueues(qm.OrderBy("sequence")).All(context.Background(), db)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	queues := make([]dbModels.DownloadQueue, 0)
	// 	downloadStatus := make([]int, 0)
	// 	for _, queuePtr := range queuesPtr {
	// 		if queuePtr.Downloading == 1 {
	// 			files, _ := ioutil.ReadDir("./.temp/" + strconv.FormatInt(queuePtr.SN, 10))
	// 			downloadStatus = append(downloadStatus, len(files))
	// 		}
	// 		queues = append(queues, *queuePtr)
	// 	}
	// 	return models.QueueStatus{Queue: queues, DownloadingStatus: downloadStatus}
	// })

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	serverMux := http.NewServeMux()

	serverMux.Handle("/", http.FileServer(pkger.Dir("/dist")))
	serverMux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.Replace(r.URL.Path, "/api", "https://api.gamer.com.tw/mobile_app", -1)
		var res *http.Response
		if r.Method == http.MethodGet {
			res, err = http.Get(url)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
		//TODO: POST method

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(res.StatusCode)
		w.Write(resBody)
	})

	go http.Serve(listener, serverMux)

	app.Navigate(fmt.Sprintf("http://%s", listener.Addr()))
	app.Run()

}
