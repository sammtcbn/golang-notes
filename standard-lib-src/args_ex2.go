package main

import (
    "fmt"
    "os"
)

func main() {
    for _, arg := range os.Args {
        fmt.Println (arg)
    }
}

/* Result:

$ go build args_ex2.go
$ ./args_ex2 11 22 33
./args_ex2
11
22
33

*/
