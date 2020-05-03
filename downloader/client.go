package downloader

import (
	"net/http"
	"net/http/cookiejar"
)

type animationDownloadClient struct {
	*http.Client
	deviceID string
	sn       string
}

func newAnimationDownloadClient(sn string) *animationDownloadClient {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}
	deviceID := getDeviceID(client)
	newAnimationDownloadClient := &animationDownloadClient{
		client,
		deviceID,
		sn,
	}
	return newAnimationDownloadClient
}
