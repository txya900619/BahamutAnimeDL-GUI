package queue

import (
	"context"
	"database/sql"
	dbModel "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/downloader"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"log"
	"strconv"
	"sync"
)

type system struct {
	db            *sql.DB
	stopper       map[string]*bool
	maxDownloader int
}

func New(db *sql.DB, maxDownloader int) *system {
	return &system{db: db, stopper: make(map[string]*bool), maxDownloader: maxDownloader}
}

func (queue *system) Start() {
	for {
		if count, _ := dbModel.DownloadQueues(qm.Where("stop=? and downloading=?", 0, 0)).Count(context.Background(), queue.db); count > 0 {
			if len(queue.stopper) < queue.maxDownloader {
				toDownload, err := dbModel.DownloadQueues(qm.Where("stop=? and downloading=?", 0, 0)).One(context.Background(), queue.db)
				if err != nil {
					log.Fatal(err)
				}
				go queue.newDownloader(strconv.Itoa(int(toDownload.SN)))
			}
		}
	}
}

func (queue *system) newDownloader(sn string) {
	stop := false
	var wg sync.WaitGroup
	queue.stopper[sn] = &stop

	intSn, err := strconv.ParseInt(sn, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	downloadingQueue, err := dbModel.FindDownloadQueue(context.Background(), queue.db, intSn)
	if err != nil {
		log.Fatal(err)
	}
	downloadingQueue.Downloading = 1
	_, err = downloadingQueue.Update(context.Background(), queue.db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go downloader.DownloadAnimation(sn, &stop, &wg)
	wg.Wait()
	delete(queue.stopper, sn)
}

func (queue *system) Stop(sn string) {
	intSn, err := strconv.ParseInt(sn, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	downloadingQueue, err := dbModel.FindDownloadQueue(context.Background(), queue.db, intSn)
	if err != nil {
		log.Fatal(err)
	}
	downloadingQueue.Stop = 1
	_, err = downloadingQueue.Update(context.Background(), queue.db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	*queue.stopper[sn] = true
}

func (queue *system) ReStart(sn string) {
	intSn, err := strconv.ParseInt(sn, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	downloadingQueue, err := dbModel.FindDownloadQueue(context.Background(), queue.db, intSn)
	if err != nil {
		log.Fatal(err)
	}
	downloadingQueue.Stop = 0
	_, err = downloadingQueue.Update(context.Background(), queue.db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}
}
