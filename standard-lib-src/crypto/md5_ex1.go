// Refer to https://blog.csdn.net/skh2015java/article/details/53468434
package main

import (
    "fmt"
    "crypto/md5"
    "io"
)

func main() {
    str := "0123456789"

//method 1
    data := []byte(str)
    has := md5.Sum (data)
    fmt.Println ("has is")
    fmt.Print (has)
    fmt.Println ("\n")

    md5str1 := fmt.Sprintf("%x", has)
    fmt.Println ("md5str1 is")
    fmt.Println (md5str1)
    fmt.Println ("\n")

//method 2
    w := md5.New()
    io.WriteString (w, str)
    fmt.Println ("w is")
    fmt.Println (w)
    fmt.Println ("\n")

    fmt.Println ("w.Sum(nil) is")
    fmt.Println (w.Sum(nil))
    fmt.Println ("\n")

    md5str2 := fmt.Sprintf("%x", w.Sum(nil))
    fmt.Println ("md5str2 is")
    fmt.Println (md5str2)
}

/*

$ echo -n 0123456789 | md5sum
781e5e245d69b566979b86e28d23f2c7  -

$ go run md5_ex1.go
has is
[120 30 94 36 93 105 181 102 151 155 134 226 141 35 242 199]

md5str1 is
781e5e245d69b566979b86e28d23f2c7


w is
&{[1732584193 4023233417 2562383102 271733878] [48 49 50 51 52 53 54 55 56 57 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] 10 10}


w.Sum(nil) is
[120 30 94 36 93 105 181 102 151 155 134 226 141 35 242 199]


md5str2 is
781e5e245d69b566979b86e28d23f2c7

*/
