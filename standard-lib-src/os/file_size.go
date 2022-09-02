package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.Stat("abc.txt")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%d bytes\n", f.Size())
    }
}
