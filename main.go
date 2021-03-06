package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	downloadHTTPContent()
}

func downloadHTTPContent() {
	downloadFileName := "random-test-file"
	downloadFileURL := "https://raw.githubusercontent.com/complexorganizations/bandwidth-waster/main/random-test-file"
	startTime := time.Now()
	for {
		err := downloadFile(downloadFileName, downloadFileURL)
		if err != nil {
			log.Println(err)
		}
		os.Remove(downloadFileName)
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
