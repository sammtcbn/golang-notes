package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    fmt.Println("Location: ", t.Location())
    fmt.Println("Time: ", t)
    fmt.Println("Time: ", t.String())
    fmt.Println(t.Format("2006-01-02 15:04:05"))
    fmt.Println(t.Format("2006-01-02"))
}

/* Result:

Location:  Local
Time:  2022-07-28 00:17:13.375063135 +0800 CST m=+0.000075355
Time:  2022-07-28 00:17:13.375063135 +0800 CST m=+0.000075355
2022-07-28 00:17:13
2022-07-28

*/
