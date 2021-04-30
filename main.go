package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	startTime        = time.Now()
	downloadFileName = randomString(64)
	downloadFileURL  = "https://raw.githubusercontent.com/complexorganizations/bandwidth-waster/main/random-test-file"
	downloadFlag     bool
	uploadFlag       bool
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
}

func main() {
	// Download the file
	if downloadFlag {
		downloadHTTPContent()
	} else if uploadFlag {
		uploadHTTPContent()
	} else {
		os.Exit(0)
	}
}

func uploadHTTPContent() {
	// If the file does not exist, you must first download it before uploading it.
	if !fileExists(downloadFileName) {
		err := downloadFile(downloadFileName, downloadFileURL)
		if err != nil {
			log.Println(err)
		}
	}
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
			fmt.Println(time.Since(startTime), "Time Running")
		}
	}
}

// Download the files to your hard drive and then delete them.
func downloadHTTPContent() {
	for {
		err := downloadFile(downloadFileName, downloadFileURL)
		if err != nil {
			log.Println(err)
		}
		err = os.Remove(downloadFileName)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(time.Since(startTime), "Time Running")
	}
}

// Download the file to the system
func downloadFile(filepath, url string) error {
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
