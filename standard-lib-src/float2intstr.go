package main

import (
    "fmt"
    "math"
    "strconv"
)

func FloatToIntStr (input_num float64) string {
    var str string
    var v float64 = math.Trunc(input_num)
    str = strconv.FormatFloat(v, 'f', 0, 64)
    return str
}

func main() {
    var floatvalue float64 = 3.14159265358979
    var str string = FloatToIntStr (floatvalue)
    fmt.Println(str)

    floatvalue = 12.99999
    str = FloatToIntStr (floatvalue)
    fmt.Println(str)
}

/* Result:

$ go run float2intstr.go
3
12

*/
