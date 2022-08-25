package main

import (
    "fmt"
)

func main() {
    var str1 string = "123"
    fmt.Println (str1)

    var str2 string = `456`
    fmt.Println (str2)

    var str3 string = `789 
  abc`
    fmt.Println (str3)

/* below will have error
    var str4 string = "aaa
   bbb"

error message:
   syntax error: unexpected literal " at end of statement

*/

}

/* Result:

$ go run str1.go
123
456
789
  abc

*/
