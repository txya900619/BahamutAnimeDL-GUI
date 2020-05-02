//+build !dev

package main

import (
	"fmt"
	"github.com/markbates/pkger"
	"github.com/txya900619/BahamutAnimeDL-GUI/crawler"
	"github.com/txya900619/BahamutAnimeDL-GUI/models"
	"github.com/txya900619/BahamutAnimeDL-GUI/utilities"
	"github.com/zserge/webview"
	"log"
	"net"
	"net/http"
	"strings"
)

var NewAnimeList []models.NewAnime
var AnimeList []models.Anime

func init() {
	NewAnimeList = crawler.GetNewAnimeList()
	go func() {
		AnimeList = crawler.GetAllAnimeList()
	}()
}

func main() {
	debug:=true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("BahaDL")
	w.SetSize(1200,800,webview.HintNone)


	w.Bind("getNewAnimeList", func() string {
		return utilities.ToJson(NewAnimeList)
	})
	w.Bind("getAllAnimeList", func() string {
		return utilities.ToJson(AnimeList)
	})
	w.Bind("getAnimesByPage", func(page int) string {
		return utilities.ToJson(AnimeList[(page-1)*18 : page*18])
	})
	w.Bind("getMaxPage", func() int {
		return len(AnimeList)/18 + 1
	})

	w.Bind("getAnimesByFilter", func(filter string) string {
		filteredAnimes := make([]models.Anime, 0)
		filter = strings.ToLower(filter)
		for _, v := range AnimeList {
			if strings.Contains(strings.ToLower(v.Title), filter) {
				filteredAnimes = append(filteredAnimes, v)
			}
		}
		return utilities.ToJson(filteredAnimes)
	})

	w.Bind("getRealSn", func(ref string) string {
		return crawler.GetRealSn(ref)
	})

	w.Bind("getAnimeAllSn", func(sn string) string {
		return utilities.ToJson(crawler.GetSnsByOneSn(sn))
	})

	net, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer net.Close()

	go http.Serve(net, http.FileServer(pkger.Dir("/dist")))
	w.Navigate(fmt.Sprintf("http://%s", net.Addr()))
	w.Run()
}
