package main

import "fmt"
import "strings"

func main() {
    var str string = "2023/08/21"
    str2 := strings.Replace(str, "/", "-" ,2)
    fmt.Println (str2)
}

/* Result:

$ go run str_replace.go
2023-08-21

*/
