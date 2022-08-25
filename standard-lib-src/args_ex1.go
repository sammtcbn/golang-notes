package main

import (
    "fmt"
    "os"
)

func main() {
    var argc int = len (os.Args)
    fmt.Println ("argc is ", argc)

    for i:=0 ; i<argc ; i++ {
        fmt.Printf ("argv[%d] is %s\n", i, os.Args[i])
    }
}

/* Result:

$ go build args_ex1.go
$ ./args_ex1 1 23 456 abc d ef
argc is  7
argv[0] is ./args_ex
argv[1] is 1
argv[2] is 23
argv[3] is 456
argv[4] is abc
argv[5] is d
argv[6] is ef

*/
