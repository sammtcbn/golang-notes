package main

import (
    "log"
    "fmt"
    "github.com/idoubi/goz"
)

func main() {
    cli := goz.NewClient()

	resp, err := cli.Get("https://finance.yahoo.com/quote/2330.TW")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", resp)
}
