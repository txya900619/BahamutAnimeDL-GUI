package database

import (
	"context"
	"database/sql"
	dbModels "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"strconv"
)

func DeleteQueue(db *sql.DB, sn string) error {

	intSn, err := strconv.ParseInt(sn, 10, 64)
	if err != nil {
		return err
	}

	findQueue, err := dbModels.FindDownloadQueue(context.Background(), db, intSn)
	if err != nil {
		return err
	}
	sequence := findQueue.Sequence
	_, err = findQueue.Delete(context.Background(), db)
	if err != nil {
		return err
	}
	if count, _ := dbModels.DownloadQueues().Count(context.Background(), db); count >= 1 {
		queues, err := dbModels.DownloadQueues().All(context.Background(), db)
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

	findQueue, err := dbModels.FindDownloadQueue(context.Background(), db, intSn)
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
