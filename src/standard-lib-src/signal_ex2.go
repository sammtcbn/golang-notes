// refer to https://cloud.tencent.com/developer/article/1645996
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    ch := make(chan os.Signal)

    signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

    go func() {
        for s := range ch {
            switch s {
                case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
                    fmt.Println("Program Exit...", s)
                    GracefullExit()
                case syscall.SIGUSR1:
                    fmt.Println("usr1 signal", s)
                case syscall.SIGUSR2:
                    fmt.Println("usr2 signal", s)
                default:
                    fmt.Println("other signal", s)
            }
        }
    }()

    fmt.Println("Program Start...")

    sum := 0
    for {
        sum++
        fmt.Println("sum:", sum)
        time.Sleep(time.Second)
    }

    fmt.Println("bye")
}

func GracefullExit() {
    fmt.Println("Start Exit...")
    fmt.Println("Execute Clean...")
    fmt.Println("End Exit...")
    os.Exit(0)
}

/* try command:
killall -SIGINT signal_ex2
killall -SIGTERM signal_ex2
killall -SIGUSR1 signal_ex2
killall -SIGUSR2 signal_ex2
*/