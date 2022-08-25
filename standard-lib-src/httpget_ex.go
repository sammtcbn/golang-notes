package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    resp, err := http.Get ("https://news.un.org/feed/subscribe/en/news/all/rss.xml")
    if err != nil {
        fmt.Println (err)
        os.Exit(1)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println (err)
        os.Exit(1)
    }

    fmt.Println(string(body))
}
