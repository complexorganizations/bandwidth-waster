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
		downloadURLPath = "http://speedtest-nyc1.digitalocean.com/1gb.test"
	} else if currentHour == 5 {
		downloadURLPath = "http://speedtest-ams2.digitalocean.com/1gb.test"
	} else if currentHour == 6 {
		downloadURLPath = "http://speedtest-sfo1.digitalocean.com/1gb.test"
	} else if currentHour == 7 {
		downloadURLPath = "http://speedtest-sgp1.digitalocean.com/1gb.test"
	} else if currentHour == 8 {
		downloadURLPath = "http://speedtest-lon1.digitalocean.com/1gb.test"
	} else if currentHour == 9 {
		downloadURLPath = "https://ga-us-ping.vultr.com/vultr.com.1000MB.bin"
	} else if currentHour == 10 {
		downloadURLPath = "https://il-us-ping.vultr.com/vultr.com.1000MB.bin"
	} else if currentHour == 11 {
		downloadURLPath = "https://tx-us-ping.vultr.com/vultr.com.1000MB.bin"
	} else if currentHour == 12 {
		downloadURLPath = "https://lax-ca-us-ping.vultr.com/vultr.com.1000MB.bin"
	} else if currentHour == 13 {
		downloadURLPath = "https://fl-us-ping.vultr.com/vultr.com.1000MB.bin"
	} else if currentHour == 14 {
		downloadURLPath = "https://ams-nl-ping.vultr.com/vultr.com.1000MB.bin"
	} else if currentHour == 15 {
		downloadURLPath = "https://tor-ca-ping.vultr.com/vultr.com.1000MB.bin"
	} else if currentHour == 16 {
		downloadURLPath = "https://sjo-ca-us-ping.vultr.com/vultr.com.1000MB.bin"
	} else if currentHour == 17 {
		downloadURLPath = "https://lax-ca-us-ping.vultr.com/vultr.com.1000MB.bin"
	} else if currentHour == 18 {
		downloadURLPath = "http://speedtest-sfo2.digitalocean.com/1gb.test"
	} else if currentHour == 19 {
		downloadURLPath = "http://speedtest-nyc3.digitalocean.com/1gb.test"
	} else if currentHour == 20 {
		downloadURLPath = "http://speedtest-blr1.digitalocean.com/1gb.test"
	} else if currentHour == 21 {
		downloadURLPath = "http://speedtest-sfo3.digitalocean.com/1gb.test"
	} else if currentHour == 22 {
		downloadURLPath = "http://speedtest-tor1.digitalocean.com/1gb.test"
	} else if currentHour == 23 {
		downloadURLPath = "https://speed.hetzner.de/1GB.bin"
	} else if currentHour == 24 {
		downloadURLPath = "https://syd-au-ping.vultr.com/vultr.com.1000MB.bin"
	}
	return downloadURLPath
}
