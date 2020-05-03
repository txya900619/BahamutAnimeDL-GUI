package downloader

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getDeviceID(client *http.Client) (deviceID string) {
	resp, err := client.Get("https://ani.gamer.com.tw/ajax/getdeviceid.php")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	deviceID = strings.Split(strings.Split(string(body), ":\"")[1], "\"")[0]

	return
}

func (client *animationDownloadClient) accessAD() {
	if vip, _ := checkADStatus(client); !vip {
		seeAD(client)
	}
}

func checkADStatus(client *animationDownloadClient) (vip, seenAD bool) {
	resp, err := client.Get("https://ani.gamer.com.tw/ajax/token.php?adID=0&sn=" + client.sn + "&device=" + client.deviceID + "&hash=" + randomHash())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	vip, err = strconv.ParseBool(strings.Split(strings.Split(string(body), "vip\":")[1], ",")[0])
	if err != nil {
		log.Fatal(err)
	}

	seenADint, err := strconv.Atoi(strings.Split(strings.Split(string(body), "time\":")[1], ",")[0])
	seenAD = seenADint != 0

	return
}

func seeAD(client *animationDownloadClient) {
	resp, _ := client.Get("https://ani.gamer.com.tw/ajax/videoCastcishu.php?sn=" + client.sn + "&s=194699")
	resp.Body.Close()

	time.Sleep(8 * time.Second)

	resp, _ = client.Get("https://ani.gamer.com.tw/ajax/videoCastcishu.php?sn=" + client.sn + "&s=194699&ad=end")
	resp.Body.Close()

	if _, seenAD := checkADStatus(client); !seenAD {
		seeAD(client)
	}
}
