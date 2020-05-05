package downloader

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	dbModel "github.com/txya900619/BahamutAnimeDL-GUI/database/models"
	"github.com/volatiletech/sqlboiler/boil"
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

func DownloadAnimation(sn int, stop *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	resolution := "720"
	maxThreads := 32
	animationDLClient := newAnimationDownloadClient(strconv.Itoa(sn), stop)

	err := animationDLClient.accessAD()
	if err != nil {
		if err.Error() == "stopped" {
			return
		}
		log.Fatal(err)
	}

	chunkUrls, key, err := animationDLClient.getAnimationChunkUrlsAndKey(resolution)
	if err != nil {
		if err.Error() == "stopped" {
			return
		}
		log.Fatal(err)
	}

	err = animationDLClient.concurrentDownloadAnimationChunk(chunkUrls, key, maxThreads)
	if err != nil {
		if err.Error() == "stopped" {
			return
		}
		log.Fatal(err)
	}

	err = animationDLClient.combineChunk(chunkUrls)
	if err != nil {
		if err.Error() == "stopped" {
			return
		}
		log.Fatal(err)
	}
}

func (client *animationDownloadClient) combineChunk(chunkUrls []string) error {
	if *client.stop {
		err := errors.New("stopped")
		return err
	}

	mergedFile, err := os.Create("./.temp/" + client.sn + "/main.ts")
	if err != nil {
		return err
	}

	for _, chunkUrl := range chunkUrls {
		if *client.stop {
			mergedFile.Close()
			rmMainDotTs(client.sn)
			err = errors.New("stopped")
			return err
		}

		chunkName := strings.Split(path.Base(chunkUrl), "?")[0]
		animationChunk, err := ioutil.ReadFile("./.temp/" + client.sn + "/" + chunkName)
		if err != nil {
			mergedFile.Close()
			rmMainDotTs(client.sn)
			return err
		}
		if _, err := mergedFile.Write(animationChunk); err != nil {
			mergedFile.Close()
			rmMainDotTs(client.sn)
			return err
		}
	}
	_ = mergedFile.Sync()
	mergedFile.Close()

	intSn, err := strconv.ParseInt(client.sn, 10, 64)
	if err != nil {
		rmMainDotTs(client.sn)
		return err
	}
	findQueue, err := dbModel.FindDownloadQueue(context.Background(), client.DB, intSn)
	if err != nil {
		rmMainDotTs(client.sn)
		return err
	}
	parseTsToMp4(client.sn, findQueue.Name, findQueue.Ep)

	newDownloaded := dbModel.DownloadedAnimation{SN: intSn, Title: findQueue.Name, Episode: findQueue.Ep}
	err = newDownloaded.Insert(context.Background(), client.DB, boil.Infer())
	if err != nil {
		rmMainDotTs(client.sn)
		return err
	}
	_, err = findQueue.Delete(context.Background(), client.DB)
	if err != nil {
		rmMainDotTs(client.sn)
		return err
	}
	err = os.RemoveAll("./.temp/" + client.sn)
	if err != nil {
		rmMainDotTs(client.sn)
		return err
	}

	return nil
}

func (client *animationDownloadClient) concurrentDownloadAnimationChunk(chunkUrls []string, key []byte, maxThreads int) error {
	if *client.stop {
		err := errors.New("stopped")
		return err
	}

	if _, err := os.Stat("./.temp/" + client.sn); os.IsNotExist(err) {
		err := os.Mkdir("./.temp/"+client.sn, 0777)
		if err != nil {
			return err
		}
	}

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

	var err error
	for _, url := range chunkUrls {
		if *client.stop {
			err = errors.New("stopped")
			break
		}
		ch <- url
	}
	close(ch)
	wg.Wait()
	return err
}

func downloadAnimationChunk(client *animationDownloadClient, url string, key []byte) bool {
	fileName := strings.Split(path.Base(url), "?")[0]
	fileState, err := os.Stat("./.temp/" + client.sn + "/" + fileName)
	if err == nil && fileState.Size() != 0 {
		return true
	}

	newChunk, err := os.Create("./.temp/" + client.sn + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("download err", err.Error())
		_ = newChunk.Close()
		_ = os.Remove("./.temp/" + client.sn + "/" + fileName)
		time.Sleep(500 * time.Millisecond)
		return false
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body, err = decryptAES128(body, key)
	if err != nil {
		fmt.Println("decrypt err")
		_ = newChunk.Close()
		_ = os.Remove("./.temp/" + client.sn + "/" + fileName)
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
		fmt.Println("copy err")
		_ = newChunk.Close()
		_ = os.Remove("./.temp/" + client.sn + "/" + fileName)
		time.Sleep(500 * time.Millisecond)
		return false
	}

	_ = newChunk.Close()
	return true
}
