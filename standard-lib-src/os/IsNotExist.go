package main

import (
    "fmt"
    "os"
)

func main() {
    fn := "/etc/123.txt"

    _, err := os.Stat (fn)
    if err != nil {
        fmt.Println(err)
        if os.IsNotExist(err) {
            fmt.Println (fn + " does not exist")
        }
    }

    _, err2 := os.Open (fn)
    if err2 != nil {
        fmt.Println(err2)
        if os.IsNotExist(err) {
            fmt.Println (fn + " does not exist")
        }
    }
}

/* Result:

$ go run os_IsNotExist.go
stat /etc/123.txt: no such file or directory
/etc/123.txt does not exist
open /etc/123.txt: no such file or directory
/etc/123.txt does not exist

*/
