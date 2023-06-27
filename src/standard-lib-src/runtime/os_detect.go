package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println (runtime.GOOS)

    if runtime.GOOS == "windows" {
        fmt.Println("You are running on Windows")
    } else if runtime.GOOS == "linux" {
        fmt.Println("You are running on Linux")
    } else if runtime.GOOS == "darwin" {
        fmt.Println("You are running on MAC")
    } else {
        fmt.Println("You are running on other OS")
    }

    os := runtime.GOOS
    switch os {
        case "windows":
            fmt.Println("You are running on Windows")
        case "linux":
            fmt.Println("You are running on Linux")
        case "darwin":
            fmt.Println("You are running on MAC")
        default:
            fmt.Println("You are running on other OS")
    }
}
