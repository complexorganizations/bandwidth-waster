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

var (
	arguments        = os.Args[1]
	downloadFileName = randomString(1024)
	downloadFileURL  = "https://raw.githubusercontent.com/complexorganizations/bandwidth-waster/main/random-test-file"
)

func main() {
	chooseUploadORDownload()
}

func chooseUploadORDownload() {
	switch imageFilePath {
	case "--download":
		downloadHTTPContent()
	case "--upload":
		uploadHTTPContent()
	default:
		fmt.Println("Error: format not supported")
	}
}

func uploadHTTPContent() {
	for {
		if !fileExists(downloadFileName) {
			err := downloadFile(downloadFileName, downloadFileURL)
			if err != nil {
				log.Println(err)
			}
		} else if fileExists(downloadFileName) {
			file, err := os.Open(downloadFileName)
			if err != nil {
				log.Println(err)
			}
			defer file.Close()
			req, err := http.NewRequest("POST", "https://bashupload.com/", file)
			if err != nil {
				log.Println(err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()
		}
	}
}

func downloadHTTPContent() {
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

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func randomString(bytesSize int) string {
	randomBytes := make([]byte, bytesSize)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}
