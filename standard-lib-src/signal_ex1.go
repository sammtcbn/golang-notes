package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func cleanup() {
    fmt.Println("cleanup")
}

func main() {
    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        cleanup()
        fmt.Println("exit");
        os.Exit(1)
    }()

    for {
        fmt.Println("sleeping...")
        time.Sleep(10 * time.Second)
    }

    fmt.Println("bye");
}
