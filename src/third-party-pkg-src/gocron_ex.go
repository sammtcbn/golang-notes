package main

import (
	"fmt"
	"time"
	"github.com/jasonlvhit/gocron"
)

func task1() {
	fmt.Println("I am running task1.", time.Now().Format("2006-01-02 15:04:05"))
}

func task2() {
	fmt.Println("I am running task2.", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	gocron.Every(1).Second().Do(task1)
	gocron.Every(3).Seconds().Do(task2)
	<-gocron.Start()
}

/* Result:

$ go run gocron_ex.go
I am running task1. 2020-03-19 01:58:38
I am running task1. 2020-03-19 01:58:39
I am running task1. 2020-03-19 01:58:40
I am running task2. 2020-03-19 01:58:40
I am running task1. 2020-03-19 01:58:41
I am running task1. 2020-03-19 01:58:42
I am running task1. 2020-03-19 01:58:43
I am running task2. 2020-03-19 01:58:43
I am running task1. 2020-03-19 01:58:44
I am running task1. 2020-03-19 01:58:45

*/
