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
)

var (
	downloadFlag bool
	uploadFlag   bool
	wg           sync.WaitGroup
)

func init() {
	// make sure user has passed some arguments.
	if len(os.Args) > 1 {
		// Supported Flags
		tempDownloadFlag := flag.Bool("download", false, "Download a huge number of files, then delete them.")
		tempUploadFlag := flag.Bool("upload", false, "Just to strees the network, upload random big files.")
		flag.Parse()
		downloadFlag = *tempDownloadFlag
		uploadFlag = *tempUploadFlag
	} else {
		log.Fatal("Error: There are no guidelines for what to do")
	}
	// Need to choose what to do
	if !uploadFlag && !downloadFlag {
		log.Fatal("Error: It is necessary for you to select whether you want to download or upload the file.")
	}
	// Cant do both at the same time.
	if uploadFlag && downloadFlag {
		log.Fatal("Error: You can't upload and download files at the same time.")
	}
}

func main() {
	if downloadFlag {
		downloadHTTPContent()
	}
	if uploadFlag {
		for {
			wg.Add(1)
			go uploadHTTPContent()
		}
	}
}

func uploadHTTPContent() {
	localTempFile := ".delete"
	// open the file and if its not there create one.
	filePath, err := os.OpenFile(localTempFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	// write the content to the file
	_, err = filePath.WriteString(string(randomString(524288)))
	if err != nil {
		log.Println(err)
	}
	// close the file
	filePath.Close()
	file, err := os.Open(localTempFile)
	if err != nil {
		log.Println(err)
	}
	file.Close()
	req, err := http.NewRequest("POST", "https://bashupload.com/", file)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	wg.Done()
}

// Download the files to your hard drive and then delete them.
func downloadHTTPContent() {
	downloadFileURL := "https://raw.githubusercontent.com/complexorganizations/bandwidth-waster/main/random-test-file"
	for {
		wg.Add(1)
		go downloadFile(downloadFileURL)
	}
}

// Download the file to the system
func downloadFile(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Error: When attempting to send your request, an error occurred.")
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	_ = body
	wg.Done()
}

// Generate a random string
func randomString(bytesSize int) string {
	randomBytes := make([]byte, bytesSize)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}
