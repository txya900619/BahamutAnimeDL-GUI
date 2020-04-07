//+build dev

package main

import (
	"fmt"
	"github.com/txya900619/BahamutAnimeDL-GUI/crawler"
	"github.com/txya900619/BahamutAnimeDL-GUI/model"
	"github.com/txya900619/BahamutAnimeDL-GUI/utilities"
	"github.com/zserge/lorca"
	"log"
	"runtime"
)

var NewAnimeList []model.NewAnime
var AnimeList []model.Anime

func init() {
	NewAnimeList = crawler.GetNewAnimeList()
	go func() {
		AnimeList = crawler.GetAllAnimeList()
	}()
}

func main() {

	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}

	app, err := lorca.New("", "", 1200, 800, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	app.Bind("getNewAnimeList", func() string {
		return utilities.ToJson(NewAnimeList)
	})
	app.Bind("getAllAnimeList", func() string {
		return utilities.ToJson(AnimeList)
	})

	app.Load(fmt.Sprintf("http://%s", "127.0.0.1:8080"))
	<-app.Done()
}
