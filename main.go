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
	downloadFileName = randomString(64)
	downloadFileURL  = "https://raw.githubusercontent.com/complexorganizations/bandwidth-waster/main/random-test-file"
	downloadFlag     bool
	uploadFlag       bool
	wg               sync.WaitGroup
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
	} else if uploadFlag {
		uploadHTTPContent()
	}
}

func uploadHTTPContent() {
	// If the file exists than start a loop of uploading it.
	if fileExists(downloadFileName) {
		for {
			file, err := os.Open(downloadFileName)
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
		}
	}
}

// Download the files to your hard drive and then delete them.
func downloadHTTPContent() {
	for {
		wg.Add(1)
		go downloadFile(downloadFileName, downloadFileURL)
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
	localFilePath := ".delete"
	os.Remove(localFilePath)
	// open the file and if its not there create one.
	filePath, err := os.OpenFile(localFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	// write the content to the file
	_, err = filePath.WriteString(string(body))
	if err != nil {
		log.Println(err)
	}
	// close the file
	filePath.Close()
	response.Body.Close()
	wg.Done()
}

// Check if the file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
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
