package main

import (
    "fmt"
    "os/exec"
)

func main() {
    cmd := exec.Command("date")
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("run fail: %s\n", err)
    }
    fmt.Printf("combined out:\n%s\n", string(out))
}

/* Result:

$ LANG=en go run exec_ex.go
combined out:
Fri Mar 20 21:12:14 CST 2020

*/
