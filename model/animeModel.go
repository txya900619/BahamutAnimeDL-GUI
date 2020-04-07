package model

type AnimeInfo struct {
	Title string `json.go:"title"`
	Img   string `json.go:"img"`
}
type NewAnime struct {
	AnimeInfo
	Sn string
}
type Anime struct {
	AnimeInfo
	Ref string
}
