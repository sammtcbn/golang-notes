// refer to https://golangcode.com/while-true/
package main

import (
    "fmt"
    "time"
)

func main() {

    for {
        fmt.Println("Infinite Loop 1")
        time.Sleep(time.Second)
    }

    // Alternative Version
    for true {
        fmt.Println("Infinite Loop 2")
        time.Sleep(time.Second)
    }
}
