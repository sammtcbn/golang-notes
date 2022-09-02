package main

import (
    "fmt"
    "strings"
)

func FloatStringShrink (instr string) string {
    var dotaddr int = -1
    var outstr string = instr
    var outstrLen int = len (outstr)
    var outstrLen2 int = -1
    var endaddr int = outstrLen

    dotaddr = strings.Index (outstr, ".")

    if (dotaddr == -1) {
        return outstr
    }

    // ex: 10.
    if (dotaddr == (outstrLen - 1)) {
        return outstr[0:outstrLen-1]
    }

    for i:=outstrLen-1; i>dotaddr; i-- {
        //fmt.Println(i, string(outstr[i]))
        if (outstr[i] == '0') {
            endaddr = i
        } else {
            break
        }
    }

    // ex: 10.00 -> 10.
    outstrLen2 = len (outstr[0:endaddr])

    if (dotaddr == (outstrLen2 - 1)) {
        endaddr = dotaddr
    }

    return outstr[0:endaddr]
}

func main() {
    var mystr string

    strs := []string {
        "12.34",
        "00.900",
        "0.80",
        "400.503",
        "103",
        "13.",
        "1.00",
        "0.00",
    }

    for _, e:= range strs {
        fmt.Println ("instr  =", e)
        mystr = FloatStringShrink (e)
        fmt.Println ("outstr =", mystr, "\n")
    }
}

/* Result:

$ go run float_string_shrink.go
instr  = 12.34
outstr = 12.34

instr  = 00.900
outstr = 00.9

instr  = 0.80
outstr = 0.8

instr  = 400.503
outstr = 400.503

instr  = 103
outstr = 103

instr  = 13.
outstr = 13

instr  = 1.00
outstr = 1

instr  = 0.00
outstr = 0

*/
