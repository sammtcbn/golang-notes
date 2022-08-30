package main

import (
    "fmt"
    "html"
)

func main() {
    var str1 string = "S&amp;P500"
    var str2 string
    str2 = html.UnescapeString(str1)
    fmt.Println (str2)
}

/* Result:

$ go run html_remove_escape_string.go
S&P500

*/
