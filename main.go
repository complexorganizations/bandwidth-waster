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

func changeFileViaHour() string {
	currentHour := time.Now().Hour()
	var downloadURLPath string
	if currentHour == 1 {
		downloadURLPath = "https://sgp-ping.vultr.com/vultr.com.1000MB.bin"
	} else if currentHour == 2 {
		downloadURLPath = "https://speed.hetzner.de/1GB.bin"
	} else if currentHour == 3 {
		downloadURLPath = "https://fastest.fish/lib/downloads/1GiB.bin"
	} else if currentHour == 4 {
		downloadURLPath = ""
	} else if currentHour == 5 {
		downloadURLPath = ""
	} else if currentHour == 6 {
		downloadURLPath = ""
	} else if currentHour == 7 {
		downloadURLPath = ""
	} else if currentHour == 8 {
		downloadURLPath = ""
	} else if currentHour == 9 {
		downloadURLPath = ""
	} else if currentHour == 10 {
		downloadURLPath = ""
	} else if currentHour == 11 {
		downloadURLPath = ""
	} else if currentHour == 12 {
		downloadURLPath = ""
	} else if currentHour == 13 {
		downloadURLPath = ""
	} else if currentHour == 14 {
		downloadURLPath = ""
	} else if currentHour == 15 {
		downloadURLPath = ""
	} else if currentHour == 16 {
		downloadURLPath = ""
	} else if currentHour == 17 {
		downloadURLPath = ""
	} else if currentHour == 18 {
		downloadURLPath = ""
	} else if currentHour == 19 {
		downloadURLPath = ""
	} else if currentHour == 20 {
		downloadURLPath = ""
	} else if currentHour == 21 {
		downloadURLPath = ""
	} else if currentHour == 22 {
		downloadURLPath = ""
	} else if currentHour == 23 {
		downloadURLPath = ""
	} else if currentHour == 24 {
		downloadURLPath = ""
	}
	return downloadURLPath
}

func downloadHTTPContent() {
	downloadFileName := "delete-this-file"
	startTime := time.Now()
	for {
		err := downloadFile(downloadFileName, changeFileViaHour())
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
