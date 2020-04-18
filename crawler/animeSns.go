package crawler

import (
	"log"
	"net/http"
	"strings"
)

func GetRealSn(ref string) string {
	resp, err := http.Get("https://ani.gamer.com.tw/animeRef.php?sn=" + ref)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(resp.Request.URL.RawQuery, "=")[1]
}
