package main
import "fmt"


func main() {
    var position int = 34
    fmt.Println(fibonacci(position))
}

func fibonacci(i int) int {
    if(i<2) {
        return i;
    }
    return fibonacci(i-2) + fibonacci(i-1);
}
