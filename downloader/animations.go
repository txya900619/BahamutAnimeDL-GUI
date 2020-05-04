package downloader

import (
	"bytes"
	"context"
	dbModel "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func DownloadAnimation(sn int, stop chan string) {
	resolution := "720"
	maxThreads := 32
	animationDLClient := newAnimationDownloadClient(strconv.Itoa(sn), stop)
	animationDLClient.accessAD()
	chunkUrls, key := animationDLClient.getAnimationChunkUrlsAndKey(resolution)
	animationDLClient.concurrentDownloadAnimationChunk(chunkUrls, key, maxThreads)
}

func (client *animationDownloadClient) combineChunk(chunkUrls []string) {
	mergedFile, err := os.Create("./temp/" + client.sn + "/main.ts")
	if err != nil {
		log.Fatal(err)
	}
	defer mergedFile.Close()
	for _, chunkUrl := range chunkUrls {
		chunkName := strings.Split(path.Base(chunkUrl), "?")[0]
		animationChunk, err := ioutil.ReadFile("./temp/" + client.sn + "/" + chunkName)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := mergedFile.Write(animationChunk); err != nil {
			log.Fatal(err)
		}
	}
	_ = mergedFile.Sync()

	intSn, err := strconv.ParseInt(client.sn, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	findQueue, err := dbModel.FindDownloadQueue(context.Background(), client.DB, intSn)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat("download"); os.IsNotExist(err) {
		err := os.Mkdir("download", os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat(findQueue.Name); os.IsNotExist(err) {
		err := os.Mkdir(findQueue.Name, os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}
	}
	savePath := "./download/" + findQueue.Name + "/" + findQueue.Name + "[" + findQueue.Ep + "]" + ".mp4"

	err = exec.Command("ffmpeg", "-y", "-i", "./temp/"+client.sn+"/main.ts", "-c", "copy", savePath).Run()
	if err != nil {
		if runtime.GOOS == "windows" {
			_ = exec.Command("./ffmpeg.exe", "-y", "-i", "./temp/"+client.sn+"/main.ts", "-c", "copy", savePath).Run()
		} else {
			_ = exec.Command("./ffmpeg", "-y", "-i", "./temp/"+client.sn+"/main.ts", "-c", "copy", savePath).Run()
		}
	}
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
