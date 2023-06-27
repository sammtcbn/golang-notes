package main

import (
    "fmt"
    "os"
    "log"
    "io/ioutil"
)

func main() {
    file, err := os.Open("read_file1.txt")
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    defer file.Close()

    content, err := ioutil.ReadAll(file)
    fmt.Print(string(content))
}

/* Result:

read_file1.txt is UTF-8 file.

$ cat read_file1.txt
1234
5678

$ go run read_file1.go
1234
5678

*/
