package database

import (
	"context"
	"database/sql"
	dbModel "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"github.com/volatiletech/sqlboiler/boil"
	"strconv"
)

func DeleteQueue(db *sql.DB, sn string) error {

	intSn, err := strconv.ParseInt(sn, 10, 64)
	if err != nil {
		return err
	}

	findQueue, err := dbModel.FindDownloadQueue(context.Background(), db, intSn)
	if err != nil {
		return err
	}
	sequence := findQueue.Sequence
	_, err = findQueue.Delete(context.Background(), db)
	if err != nil {
		return err
	}
	if count, _ := dbModel.DownloadQueues().Count(context.Background(), db); count >= 1 {
		queues, err := dbModel.DownloadQueues().All(context.Background(), db)
		if err != nil {
			return err
		}
		for _, queue := range queues {
			if queue.Sequence > sequence {
				queue.Sequence -= 1
				_, err = queue.Update(context.Background(), db, boil.Infer())
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func QueueDisDownloading(db *sql.DB, sn string) error {
	intSn, err := strconv.ParseInt(sn, 10, 64)
	if err != nil {
		return err
	}

	findQueue, err := dbModel.FindDownloadQueue(context.Background(), db, intSn)
	if err != nil {
		return err
	}
	findQueue.Downloading = 0
	_, err = findQueue.Update(context.Background(), db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
