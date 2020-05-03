package downloader

import (
	"strconv"
)

func DownloadAnimation(sn int) {
	animationDLClient := newAnimationDownloadClient(strconv.Itoa(sn))
	animationDLClient.accessAD()
}
