package main

import "fmt"
import "strings"

func main() {
    x := "123456789abcdefg"
    fmt.Println("Str: ", x);
    i := strings.Index(x, "a")
    fmt.Println("index of a: ", i)

    if i > -1 {
        part1 := x[:i]
        part2 := x[i+1:]
        fmt.Println("part1: ", part1)
        fmt.Println("part2: ", part2)
    } else {
        fmt.Println("Index not found")
        fmt.Println(x)
    }
}

/* Result:

$ go run find_substring2.go
Str:  123456789abcdefg
index of a:  9
part1:  123456789
part2:  bcdefg

*/
