package downloader

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

func DownloadAnimation(sn int) {
	resolution := "720"
	maxThreads := 32
	animationDLClient := newAnimationDownloadClient(strconv.Itoa(sn))
	animationDLClient.accessAD()
	chunkUrls, key := animationDLClient.getAnimationChunkUrlsAndKey(resolution)
	animationDLClient.concurrentDownloadAnimationChunk(chunkUrls, key, maxThreads)

}

func (client *animationDownloadClient) concurrentDownloadAnimationChunk(chunkUrls []string, key []byte, maxThreads int) {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(maxThreads)

	for i := 0; i < maxThreads; i++ {
		go func() {
			for {
				url, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				for {
					if downloadAnimationChunk(client, url, key) {
						break
					}
				}
			}
		}()
	}

	for _, url := range chunkUrls {
		ch <- url
	}
	close(ch)
	wg.Wait()
}

func downloadAnimationChunk(client *animationDownloadClient, url string, key []byte) (success bool) {
	fileName := strings.Split(path.Base(url), "?")[0]
	fileState, err := os.Stat("./.temp/" + fileName)
	if err == nil && fileState.Size() != 0 {
		return true
	}

	newChunk, err := os.Create("./.temp/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Get(url)
	if err != nil {
		_ = newChunk.Close()
		_ = os.Remove("./.temp/" + fileName)
		time.Sleep(500 * time.Millisecond)
		return false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_ = newChunk.Close()
		_ = os.Remove("./.temp/" + fileName)
		time.Sleep(500 * time.Millisecond)
		return false
	}

	body, err = decryptAES128(body, key)
	if err != nil {
		_ = newChunk.Close()
		_ = os.Remove("./.temp/" + fileName)
		time.Sleep(500 * time.Millisecond)
		return false
	}

	syncByte := uint8(71) //0x47
	bodyLen := len(body)
	for j := 0; j < bodyLen; j++ {
		if body[j] == syncByte {
			body = body[j:]
			break
		}
	}

	_, err = io.Copy(newChunk, bytes.NewReader(body))
	if err != nil {
		_ = newChunk.Close()
		_ = os.Remove("./.temp/" + fileName)
		time.Sleep(500 * time.Millisecond)
		return false
	}

	_ = newChunk.Close()
	return true
}
