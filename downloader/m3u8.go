package downloader

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/grafov/m3u8"
	dbModels "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"io/ioutil"
	"strconv"
	"strings"
)

func (client *animationDownloadClient) getAnimationChunkUrlsAndKey(resolution string) ([]string, []byte, error) {
	if *client.stop {
		err := errors.New("stopped")
		return nil, nil, err
	}

	chunksListUrl, err := findChunkListMatchResolution(client, resolution)
	if err != nil {
		return nil, nil, err
	}
	resp, err := client.Get(chunksListUrl)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var key []byte
	chunkUrls := make([]string, 0)
	playList, listType, err := m3u8.Decode(*bytes.NewBuffer(body), true)
	if err != nil {
		return nil, nil, err
	}
	if listType == m3u8.MEDIA {
		mediaPlayList := playList.(*m3u8.MediaPlaylist)

		resp, err := client.Get(mediaPlayList.Key.URI)
		if err != nil {
			return nil, nil, err
		}
		defer resp.Body.Close()
		key, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, err
		}

		chunkUrlsPrefix := strings.Split(chunksListUrl, "chunklist")[0]
		for _, chunk := range mediaPlayList.Segments {
			if chunk != nil {
				chunkUrls = append(chunkUrls, chunkUrlsPrefix+chunk.URI)
			}
		}
	}
	intSn, _ := strconv.ParseInt(client.sn, 10, 64)
	downloadQueue, err := dbModels.FindDownloadQueue(context.Background(), client.DB, intSn)
	if err != nil {
		return nil, nil, err
	}
	downloadQueue.NumberOfChunk = int64(len(chunksListUrl))
	_, err = downloadQueue.Update(context.Background(), client.DB, boil.Infer())
	if err != nil {
		return nil, nil, err
	}

	return chunkUrls, key, nil
}

func findChunkListMatchResolution(client *animationDownloadClient, resolution string) (string, error) {
	if *client.stop {
		err := errors.New("stooped")
		return "", err
	}

	m3u8Url, err := getAnimationM3u8Url(client)
	if err != nil {
		return "", err
	}
	resp, err := client.Get(m3u8Url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	chunkListUrl := ""
	playList, listType, err := m3u8.Decode(*bytes.NewBuffer(body), true)
	if err != nil {
		return "", err
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

	return chunkListUrl, nil
}

func getAnimationM3u8Url(client *animationDownloadClient) (string, error) {
	if *client.stop {
		err := errors.New("stopped")
		return "", err
	}

	resp, err := client.Get("https://ani.gamer.com.tw/ajax/m3u8.php?sn=" + client.sn + "&device=" + client.deviceID)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	jsonParse := make(map[string]string)
	err = json.Unmarshal(body, &jsonParse)
	if err != nil {
		return "", err
	}
	m3u8ListUrl := "https:" + jsonParse["src"]

	return m3u8ListUrl, nil
}
