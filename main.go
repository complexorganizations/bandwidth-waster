package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

var (
	downloadFileName = "index.html"
	downloadURLPath  = "http://127.0.0.1"
)

func main() {
	hostHTTPContent()
	downloadHTTPContent()
}

func hostHTTPContent() {
	http.Handle("/", http.FileServer(http.Dir("www")))
	http.ListenAndServe(":80", nil)
}

func downloadHTTPContent() {
	for {
		err := downloadFile(downloadFileName, downloadURLPath)
		if err != nil {
			log.Println(err)
		}
		os.Remove(downloadFileName)
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
