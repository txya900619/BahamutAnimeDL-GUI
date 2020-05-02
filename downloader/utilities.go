package downloader

import (
	"math/rand"
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
