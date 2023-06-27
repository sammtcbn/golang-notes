package main

import "fmt"
import "strings"

func indexAt(content string, sep string, startaddr int) int {
    idx := strings.Index(content[startaddr:], sep)
    if idx > -1 {
        idx += startaddr
    }
    return idx
}

func main() {
    x := "123456789abcdefg"
    fmt.Println("Str: ", x);

    i := indexAt(x, "123", 0)
    fmt.Println("index (from started) of 123 (search from index 0):", i)

    i = indexAt(x, "456", 0)
    fmt.Println("index (from started) of 456 (search from index 0):", i)

    i = indexAt(x, "456", 3)
    fmt.Println("index (from started) of 456 (search from index 3):", i)

    i = indexAt(x, "456", 4)
    fmt.Println("index (from started) of 456 (search from index 4):", i)
}

/* Result:

$ go run indexat.go
Str:  123456789abcdefg
index (from started) of 123 (search from index 0): 0
index (from started) of 456 (search from index 0): 3
index (from started) of 456 (search from index 3): 3
index (from started) of 456 (search from index 4): -1

*/
