package downloader

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getDeviceID() (deviceID string, cookie *http.Cookie) {
	resp, err := http.Get("https://ani.gamer.com.tw/ajax/getdeviceid.php")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	deviceID = strings.Split(strings.Split(string(body), ":\"")[1], "\"")[0]

	for _, respCookie := range resp.Cookies() {
		if respCookie.Name == "nologinuser" {
			cookie = respCookie
		}
	}
	return
}

func accessAD(client *http.Client, sn int, deviceID string) {
	if vip, _ := checkADStatus(client, sn, deviceID); !vip {
		seeAD(client, sn, deviceID)
	}
}

func checkADStatus(client *http.Client, sn int, deviceID string) (vip, seenAD bool) {
	resp, err := client.Get("https://ani.gamer.com.tw/ajax/token.php?adID=0&sn=" + strconv.Itoa(sn) + "&device=" + deviceID + "&hash=" + randomHash())
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

func seeAD(client *http.Client, sn int, deviceID string) {
	resp, _ := client.Get("https://ani.gamer.com.tw/ajax/videoCastcishu.php?sn=" + strconv.Itoa(sn) + "&s=194699")
	resp.Body.Close()

	time.Sleep(8 * time.Second)

	resp, _ = client.Get("https://ani.gamer.com.tw/ajax/videoCastcishu.php?sn=" + strconv.Itoa(sn) + "&s=194699&ad=end")
	resp.Body.Close()

	if _, seenAD := checkADStatus(client, sn, deviceID); !seenAD {
		seeAD(client, sn, deviceID)
	}
}
