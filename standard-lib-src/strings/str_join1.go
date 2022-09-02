package main

import "fmt"

func main() {
    var str string = "123"
    str += "456"
    str += "789"
    fmt.Println (str)
}

/* Result:

$ go run str_join1.go
123456789

*/
