package downloader

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func randomHash() string {
	rand.Seed(time.Now().UnixNano())
	dictionary := "1234567890qwertyuiopasdfghjklzxcvbnm"
	hash := ""
	for i := 0; i < 12; i++ {
		hash = hash + string(dictionary[rand.Intn(len(dictionary))])
	}
	return hash
}

func decryptAES128(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, key[:aes.BlockSize])
	decryptedData := make([]byte, len(encrypted))
	mode.CryptBlocks(decryptedData, encrypted)
	decryptedData = pkc5Unpadding(decryptedData)

	return decryptedData, nil
}

func pkc5Unpadding(decryptedData []byte) []byte {
	dataLength := len(decryptedData)
	unPadding := int(decryptedData[dataLength-1])
	return decryptedData[:(dataLength - unPadding)]
}

func parseTsToMp4(sn string, title string, episode string, spacial int64) {
	if _, err := os.Stat("download"); os.IsNotExist(err) {
		err := os.Mkdir("download", 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	if _, err := os.Stat("./download/" + title); os.IsNotExist(err) {
		err := os.Mkdir("./download/"+title, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	savePath := "./download/" + title + "/" + title + " [" + episode + "]" + ".mp4"
	if spacial == 1 {
		savePath = "./download/" + title + "/" + title + " 特別篇 [" + episode + "]" + ".mp4"
	}

	err := exec.Command("ffmpeg", "-y", "-i", "./.temp/"+sn+"/main.ts", "-c", "copy", savePath).Run()
	if err != nil {
		if runtime.GOOS == "windows" {
			err = exec.Command("./ffmpeg.exe", "-y", "-i", "./.temp/"+sn+"/main.ts", "-c", "copy", savePath).Run()
		} else {
			err = exec.Command("./ffmpeg", "-y", "-i", "./.temp/"+sn+"/main.ts", "-c", "copy", savePath).Run()
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

func waitEightSec(complete chan<- bool, stop *bool) {
	end := make(chan struct{})
	go func() {
		for {
			select {
			case <-end:
				return
			default:
				if *stop {
					complete <- false
					close(end)
					return
				}
			}
		}
	}()
	time.Sleep(8 * time.Second)
	if *stop {
		return
	}
	close(end)
	complete <- true
}

func rmMainDotTs(sn string) {
	err := os.Remove("./.temp/" + sn + "/main.ts")
	if err != nil {
		panic(err)
	}
}
