package main

import "fmt"
import "strings"

func main() {
    x := "123456789abcdefg"
    fmt.Println("Str: ", x);

    i := strings.Index(x, "a")
    fmt.Println("index of a: ", i)

    i = strings.Index(x[i:], "c")
    fmt.Println("index of c: ", i, "(started from a)")

    i = strings.Index(x[i:], "z")
    fmt.Println("index of z: ", i, "(started from c)")
}

/* Result:

$ go run find_substring3.go
Str:  123456789abcdefg
index of a:  9
index of c:  2 (started from a)
index of z:  -1 (started from c)

*/
