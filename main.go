package main

import (
	"crypto/rand"
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
	downloadFileName := randomString(64)
	downloadFileURL := fmt.Sprintf("https://ros-static.ga/public/ros-data-waster-dummy?", randomString(512))
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

func randomString(bytesSize int) string {
	randomBytes := make([]byte, bytesSize)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}
