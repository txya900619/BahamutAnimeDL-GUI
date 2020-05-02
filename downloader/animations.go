package downloader

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func DownloadAnimation(sn int) {
	cookieJar, _ := cookiejar.New(nil)
	clinet := http.Client{Jar: cookieJar}
	deviceID, cookie := getDeviceID()
	if cookie != nil {
		clinet.Jar.SetCookies(&url.URL{Scheme: "https", Host: "ani.gamer.com.tw"}, []*http.Cookie{cookie})
	}
	accessAD(&clinet, sn, deviceID)
}
