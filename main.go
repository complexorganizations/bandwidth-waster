package main

import (
        "crypto/rand"
        "fmt"
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
        http.HandleFunc("/", helloServer)
        http.ListenAndServe(":80", nil)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, randomString(1024), r.URL.Path[1:])
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

func randomString(bytesize int) string {
        randomBytes := make([]byte, bytesize)
        rand.Read(randomBytes)
        randomString := fmt.Sprintf("%X", randomBytes)
        return randomString
}
