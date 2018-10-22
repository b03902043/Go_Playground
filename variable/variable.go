package main

import "fmt"

func zz(a int, b bool) (result int) {
    result = a
    return
}

func main() {
    a := 1
    var b int
    var c, d = 1, true
    var e int = 100
    f, g := 2, "hello"
    fmt.Println(a, b, c, d, e, f, g)
    fmt.Printf("%v %v %v", a, b, c)
}
