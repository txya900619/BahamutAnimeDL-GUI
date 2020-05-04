package downloader

import (
	"bytes"
	"encoding/json"
	"github.com/grafov/m3u8"
	"io/ioutil"
	"log"
	"strings"
)

func (client *animationDownloadClient) getAnimationChunkUrlsAndKey(resolution string) (chunkUrls []string, key []byte) {
	chunksListUrl := findChunkListMatchResolution(client, resolution)
	resp, err := client.Get(chunksListUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	playList, listType, err := m3u8.Decode(*bytes.NewBuffer(body), true)
	if err != nil {
		log.Fatal(err)
	}
	if listType == m3u8.MEDIA {
		mediaPlayList := playList.(*m3u8.MediaPlaylist)

		resp, err := client.Get(mediaPlayList.Key.URI)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		key, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		chunkUrlsPrefix := strings.Split(chunksListUrl, "chunklist")[0]
		for _, chunk := range mediaPlayList.Segments {
			if chunk != nil {
				chunkUrls = append(chunkUrls, chunkUrlsPrefix+chunk.URI)
			}
		}
	}

	return
}

func findChunkListMatchResolution(client *animationDownloadClient, resolution string) (chunkListUrl string) {
	m3u8Url := getAnimationM3u8Url(client)
	resp, err := client.Get(m3u8Url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	playList, listType, err := m3u8.Decode(*bytes.NewBuffer(body), true)
	if err != nil {
		log.Fatal(err)
	}
	if listType == m3u8.MASTER {
		masterPlayList := playList.(*m3u8.MasterPlaylist)
		for _, chunkList := range masterPlayList.Variants {
			listResolution := strings.Split(chunkList.Resolution, "x")[1]
			if listResolution == resolution {
				chunkListUrl = strings.Split(m3u8Url, "playlist.m3u8")[0] + chunkList.URI
			}
		}
	}

	return
}

func getAnimationM3u8Url(client *animationDownloadClient) (m3u8ListUrl string) {
	resp, err := client.Get("https://ani.gamer.com.tw/ajax/m3u8.php?sn=" + client.sn + "&device=" + client.deviceID)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	jsonParse := make(map[string]string)
	err = json.Unmarshal(body, &jsonParse)
	if err != nil {
		log.Fatal(err)
	}
	m3u8ListUrl = "https:" + jsonParse["src"]

	return
}
