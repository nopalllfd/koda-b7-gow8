package dua

import (
	"fmt"
	"math/rand"
	"time"
)

func WebFetcher(url string, urlChn chan string) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	urlChn <- fmt.Sprintf("Fetched: %s", url)
}
