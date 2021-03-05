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
	// 1024 MB = 1GB
	megabyteToWaste  = 1024
	downloadFileName = "random-string-file"
	downloadURLPath  = "https://raw.githubusercontent.com/complexorganizations/bandwidth-waster/main/random-test-file"
	startTime        = time.Now()
)

func main() {
	downloadHTTPContent()
}

func downloadHTTPContent() {
	for loop := 0; loop <= megabyteToWaste; loop++ {
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
