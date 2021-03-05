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
	downloadURLPath  = "https://raw.githubusercontent.com/complexorganizations/bandwidth-waster/main/random-test-file"
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
		os.Remove(downloadFileName)
		fmt.Println(loop, "Gigabyte Wasted")
		fmt.Println(time.Since(startTime), "Time Running")
	}
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
