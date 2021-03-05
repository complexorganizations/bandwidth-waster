package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	megabytesToWaste = 102400000
	downloadFileName = "index.html"
	downloadURLPath  = "https://ros-static.ga/public/ros-data-waster-dummy"
)

func main() {
	downloadHTTPContent()
}

func downloadHTTPContent() {
	for loop := 0; loop <= megabytesToWaste; loop++ {
		err := downloadFile(downloadFileName, downloadURLPath)
		if err != nil {
			log.Println(err)
		}
		os.Remove(downloadFileName)
		fmt.Println(loop, "Megabytes Wasted")
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
