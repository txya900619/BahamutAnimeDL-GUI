package downloader

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func (client *animationDownloadClient) accessAD() error {
	if *client.stop {
		err := errors.New("stopped")
		return err
	}

	vip, _, err := checkADStatus(client)
	if err != nil {
		return err
	}
	if !vip {
		err = seeAD(client)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkADStatus(client *animationDownloadClient) (bool, bool, error) {
	if *client.stop {
		err := errors.New("stopped")
		return false, false, err
	}

	resp, err := client.Get("https://ani.gamer.com.tw/ajax/token.php?adID=0&sn=" + client.sn + "&device=" + client.deviceID + "&hash=" + randomHash())
	if err != nil {
		return false, false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	vip, err := strconv.ParseBool(strings.Split(strings.Split(string(body), "vip\":")[1], ",")[0])
	if err != nil {
		return false, false, err
	}

	seenADint, err := strconv.Atoi(strings.Split(strings.Split(string(body), "time\":")[1], ",")[0])
	if err != nil {
		return false, false, err
	}
	seenAD := seenADint != 0

	return vip, seenAD, nil
}

func seeAD(client *animationDownloadClient) error {
	if *client.stop {
		err := errors.New("stopped")
		return err
	}

	resp, _ := client.Get("https://ani.gamer.com.tw/ajax/videoCastcishu.php?sn=" + client.sn + "&s=194699")
	resp.Body.Close()

	waitChan := make(chan bool)
	go waitEightSec(waitChan, client.stop)
	if !<-waitChan {
		err := errors.New("stopped")
		return err
	}
	close(waitChan)

	resp, _ = client.Get("https://ani.gamer.com.tw/ajax/videoCastcishu.php?sn=" + client.sn + "&s=194699&ad=end")
	resp.Body.Close()

	_, seenAD, err := checkADStatus(client)
	if err != nil {
		return err
	}
	if !seenAD {
		err = seeAD(client)
		if err != nil {
			return err
		}
	}

	return nil
}
