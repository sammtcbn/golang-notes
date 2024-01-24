package main

import (
    "fmt"
    "strconv"
    "reflect"
)

func main() {
    var val int = 2024
    fmt.Printf("va; type : %s\n", reflect.TypeOf(val))

    year1 := strconv.Itoa(val)
    fmt.Println(year1)
    fmt.Printf("year1 type : %s\n", reflect.TypeOf(year1))

    year2 := fmt.Sprintf("%04d", val)
    fmt.Println(year2)
    fmt.Printf("year2 type : %s\n", reflect.TypeOf(year2))
}

