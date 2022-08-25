package main

import (
    "fmt"
    "strconv"
)

func FloatToString (input_num float64, decimal_place int) string {
    return strconv.FormatFloat(input_num, 'f', decimal_place, 64)
}

func main() {
    var value float64 = 3.14159265358979
    fmt.Println(FloatToString(value, 2))
    fmt.Println(FloatToString(value, 5))
}

/* Result:

$ go run float2str.go
3.14
3.14159

*/
