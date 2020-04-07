//go:generate go run -tags generate gen.go
//+build !dev

package main

import (
	"fmt"
	"github.com/txya900619/BahamutAnimeDL-GUI/crawler"
	"github.com/txya900619/BahamutAnimeDL-GUI/model"
	"github.com/txya900619/BahamutAnimeDL-GUI/utilities"
	"github.com/zserge/lorca"
	"log"
	"net"
	"net/http"
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

	net, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer net.Close()

	go http.Serve(net, http.FileServer(FS))
	app.Load(fmt.Sprintf("http://%s", net.Addr()))
	<-app.Done()
}
