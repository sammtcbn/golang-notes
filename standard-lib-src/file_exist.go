package main

import (
    "log"
    "fmt"
    "os"
)

func main() {
    fn := "abc.txt"
    _, err := os.Stat (fn)
    if err != nil {
        fmt.Println(err)
        if os.IsNotExist(err) {
            log.Fatal (fn + " does not exist")
        }
    }

    fmt.Println (fn + " exist")
}

/* Result:

$ go run file_exist.go
stat abc.txt: no such file or directory
2022/09/02 13:25:40 abc.txt does not exist
exit status 1

*/
