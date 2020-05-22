package models

import dbModels "github.com/txya900619/BahamutAnimeDL-GUI/database/models"

type AnimeInfo struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}
type NewAnime struct {
	AnimeInfo
	Sn string `json:"sn"`
}
type Anime struct {
	AnimeInfo
	Ref string `json:"ref"`
}
type Sn struct {
	Sn          string `json:"sn"`
	Number      string `json:"number"`
	CanDownload bool   `json:"canDownload"`
}

type QueueStatus struct {
	Queue             []dbModels.DownloadQueue `json:"queue"`
	DownloadingStatus []int                    `json:"downloading_status"`
}
