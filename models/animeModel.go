package models

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
