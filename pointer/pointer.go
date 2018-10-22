package pointer

import (
    "fmt"
)

type Coordinate struct {
    X, y int
}

func Struct_pointer_demo() {
    fmt.Println(Coordinate{})
    fmt.Println(Coordinate{1, 2})
    fmt.Println(Coordinate{X: 1, y: 100})

    i := 1
    p := &i  // get address of variable i
    *p += 10 // *p = reference of i
    fmt.Println(p, *p)
    p2 := &Coordinate{1, 3}
    // when point to struct, p2.X == (*p2).X
    fmt.Println(p2, *p2, p2.X, (*p2).X)
}
