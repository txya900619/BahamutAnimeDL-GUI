package downloader

import (
	"database/sql"
	"github.com/txya900619/BahamutAnimeDL-GUI/database"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

type animationDownloadClient struct {
	*http.Client
	*sql.DB
	deviceID string
	sn       string
	stop     chan string
}

func newAnimationDownloadClient(sn string, stop chan string) (newAnimationDownloadClient *animationDownloadClient) {
	cookieJar, _ := cookiejar.New(nil)
	db := database.ConnectSqlite()
	newAnimationDownloadClient = &animationDownloadClient{
		Client: &http.Client{Jar: cookieJar},
		DB:     db,
		sn:     sn,
	}
	newAnimationDownloadClient.getDeviceID()
	return
}

func (client *animationDownloadClient) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36")
	req.Header.Add("referer", "https://ani.gamer.com.tw/animeVideo.php?sn="+client.sn)
	req.Header.Add("origin", "https://ani.gamer.com.tw")
	resp, err = client.Do(req)
	return
}

func (client *animationDownloadClient) getDeviceID() {
	resp, err := client.Get("https://ani.gamer.com.tw/ajax/getdeviceid.php")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	client.deviceID = strings.Split(strings.Split(string(body), ":\"")[1], "\"")[0]

	return
}
