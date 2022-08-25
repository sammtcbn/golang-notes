package main

import (
    "fmt"
    "strings"
)

func main() {
    var idx int = -1
    str := "Hello world, sam is programming"

    idx = strings.Index (str, "sam")
    fmt.Println (idx)

    idx = strings.Index (str[14:], "sam")
    fmt.Println (idx)
}

/* Result:

$ go run find_substring1.go
13
-1

*/
