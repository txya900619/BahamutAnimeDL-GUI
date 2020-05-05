package downloader

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func (client *animationDownloadClient) accessAD() (err error) {
	if *client.stop {
		err = errors.New("stopped")
		return
	}

	vip, _, err := checkADStatus(client)
	if err != nil {
		return
	}
	if !vip {
		err = seeAD(client)
		if err != nil {
			return
		}
	}

	return
}

func checkADStatus(client *animationDownloadClient) (vip, seenAD bool, err error) {
	if *client.stop {
		err = errors.New("stopped")
		return
	}

	resp, err := client.Get("https://ani.gamer.com.tw/ajax/token.php?adID=0&sn=" + client.sn + "&device=" + client.deviceID + "&hash=" + randomHash())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	vip, err = strconv.ParseBool(strings.Split(strings.Split(string(body), "vip\":")[1], ",")[0])
	if err != nil {
		return
	}

	seenADint, err := strconv.Atoi(strings.Split(strings.Split(string(body), "time\":")[1], ",")[0])
	seenAD = seenADint != 0

	return
}

func seeAD(client *animationDownloadClient) (err error) {
	if *client.stop {
		err = errors.New("stopped")
		return
	}

	resp, _ := client.Get("https://ani.gamer.com.tw/ajax/videoCastcishu.php?sn=" + client.sn + "&s=194699")
	resp.Body.Close()

	waitChan := make(chan bool)
	waitEightSec(waitChan, client.stop)
	if !<-waitChan {
		err = errors.New("stopped")
		return
	}

	resp, _ = client.Get("https://ani.gamer.com.tw/ajax/videoCastcishu.php?sn=" + client.sn + "&s=194699&ad=end")
	resp.Body.Close()

	_, seenAD, err := checkADStatus(client)
	if err != nil {
		return
	}
	if !seenAD {
		err = seeAD(client)
		if err != nil {
			return
		}
	}

	return
}
