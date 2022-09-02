package main

import "fmt"

func main() {
    var str string = "123"
    var fv float64 = 3.54
    str = fmt.Sprintf("%s %f", str, fv)
    fmt.Println (str)
}

/* Result:

$ go run str_join2.go
123 3.540000

*/
