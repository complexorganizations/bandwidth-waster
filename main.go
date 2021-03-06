package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	gigabyteToWaste  = 1024
	downloadFileName = "random-test-file"
	downloadURLPath  = "http://speedtest-sgp1.digitalocean.com/1gb.test"
	startTime        = time.Now()
)

func main() {
	downloadHTTPContent()
}

func downloadHTTPContent() {
	for loop := 0; loop <= gigabyteToWaste; loop++ {
		err := downloadFile(downloadFileName, downloadURLPath)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(loop, "Megabytes Wasted")
		fmt.Println(time.Since(startTime), "Time Running")
	}
	os.Remove(downloadFileName)
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}
