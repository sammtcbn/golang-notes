package main

import (
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "golang.org/x/text/encoding/traditionalchinese"
    "golang.org/x/text/transform"
)

func main() {
    var contentNew string
    file, err := os.Open("read_file2_big5.txt")
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    defer file.Close()

// method 1
    content, err := ioutil.ReadAll(file)
    fmt.Println("----- method 1 -----")
    fmt.Println(string(content))

// method 2
    big5ToUTF8 := traditionalchinese.Big5.NewDecoder()
    contentNew, _, _ = transform.String(big5ToUTF8, string(content))
    //fmt.Print(string(contentNew))
    fmt.Println("----- method 2 -----")
    fmt.Println(contentNew)
}

/* Result:

read_file2_big5.txt is Big5 file.

$ go get golang.org/x/text

$ go run read_file2_big5.go
----- method 1 -----
ڬO@
----- method 2 -----
我是一隻魚

*/
