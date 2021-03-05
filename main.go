package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// 1024 Megabytes = 1 Gigabytes
// 1048576 Megabytes = 1 Terabytes
// 1073741824 Megabytes = 1 Petabytes
// 1099511627776 Megabytes = 1 Exabytes

var (
	megabyteToWaste  = 1073741824
	downloadFileName = "random-test-file"
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
