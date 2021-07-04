package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	startTime        = time.Now()
	downloadFileName = randomString(64)
	downloadFileURL  = "https://raw.githubusercontent.com/complexorganizations/bandwidth-waster/main/random-test-file"
	downloadFlag     bool
	uploadFlag       bool
	wg               sync.WaitGroup
)

func init() {
	if len(os.Args) > 1 {
		// Supported Flags
		tempDownloadFlag := flag.Bool("download", false, "download")
		tempUploadFlag := flag.Bool("upload", false, "upload")
		flag.Parse()
		downloadFlag = *tempDownloadFlag
		uploadFlag = *tempUploadFlag
	} else {
		log.Fatal("Error: There are no guidelines for what to do")
	}
	if !uploadFlag && !downloadFlag {
		log.Fatal("Error: Not a valid response")
	}
}

func main() {
	// Download the file
	if downloadFlag {
		downloadHTTPContent()
	} else if uploadFlag {
		uploadHTTPContent()
	}
}

func uploadHTTPContent() {
	// If the file exists than start a loop of uploading it.
	if fileExists(downloadFileName) {
		for {
			file, err := os.Open(downloadFileName)
			handleErrors(err)
			file.Close()
			req, err := http.NewRequest("POST", "https://bashupload.com/", file)
			handleErrors(err)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			resp, err := http.DefaultClient.Do(req)
			handleErrors(err)
			resp.Body.Close()
		}
	}
}

// Download the files to your hard drive and then delete them.
func downloadHTTPContent() {
	for {
		wg.Add(1)
		go downloadFile(downloadFileName, downloadFileURL)
		fmt.Println(time.Since(startTime))
	}
}

// Download the file to the system
func downloadFile(filepath, url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error: When attempting to send your request, an error occurred.")
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	response.Body.Close()
	_ = body
	wg.Done()
}

// Check if the file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Generate a random string
func randomString(bytesSize int) string {
	randomBytes := make([]byte, bytesSize)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}

func handleErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}
