package main

import (
    "flag"
    "fmt"
)

var (
    h bool
    ip string
    port string
)

func init() {
    flag.BoolVar(&h, "h", false, "show help")
    flag.StringVar(&ip, "i", "127.0.0.1", "assign `ip`")
    flag.StringVar(&port, "p", "1883", "assign port")
}

func main() {
    flag.Parse()

    if h {
        flag.Usage()
    } else {
        fmt.Println ("ip is", ip)
        fmt.Println ("port is", port)
    }
}

/* Result:

$ go build flag_ex1.go
$ ./flag_ex1
ip is 127.0.0.1
port is 1883

$ ./flag_ex1 -h
Usage of ./flag_ex1:
  -h    show help
  -i ip
        assign ip (default "127.0.0.1")
  -p string
        assign port (default "1883")

$ ./flag_ex1 -i 192.168.0.1 -p 8883
ip is 192.168.0.1
port is 8883

$ ./flag_ex1 -a -b
flag provided but not defined: -a
Usage of ./flag_ex1:
  -h    show help
  -i ip
        assign ip (default "127.0.0.1")
  -p string
        assign port (default "1883")

*/
