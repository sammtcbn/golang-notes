package main

import (
    "fmt"
    "strconv"
    "os"
)

func StringToFloat (str string) float64 {
    floatvalue, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println (err)
        os.Exit (2)
    }
    return floatvalue
}

func main() {
    var mystr string = "3.14159265358979"
    var myvalue float64
    myvalue = StringToFloat (mystr)
    fmt.Println (myvalue)
}

/* Result:

$ go run string2float.go
3.14159265358979

*/
