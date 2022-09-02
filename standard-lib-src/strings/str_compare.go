package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println (strings.Compare("123", "123"))
    fmt.Println (strings.Compare("123", "456"))
    fmt.Println ("hello"=="hello")
}

/* Result:

$ go run str_compare.go
0
-1
true
*/
