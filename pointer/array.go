package pointer

import (
    "fmt"
)

func Array() {
    var ar1 [5]int
    ar2 := [5]int{6, 5, 4, 3, 2}
    fmt.Println(ar1, ar2)
}

func Slice()    {
    ar := [5]int {2, 3, 4, 5, 6}
    s1 := ar[1:3]
    fmt.Printf("before: %v\n", ar)
    s1[0] = 100000000
    fmt.Printf("after: %v\n", ar)

    var s []int
    fmt.Println(s, len(s), cap(s))
    if s == nil {
        fmt.Println("nil!")
    }

    // important! create slice with "make(, len, cap)"
    sl := make([]int, 2, 3)
    printSlice("make_5_100", sl)

    sl = append(sl, 1, 2, 3)
    printSlice("append_123", sl)
}

func IterSlice()    {
    pow := []int{1, 2, 4, 8}
    for i, v := range pow   {
        fmt.Printf("2**%d = %d\n", i, v)
    }
    for i := range pow  {
        fmt.Printf("(%d %d) ", i, pow[i])
    }
    fmt.Println("")
}

func ListOfStruct() {
    s := []struct {
        name string
        value int
    }{
        {"a", 10},
        {"b", 5},
        {"c", 3},
    }
    fmt.Println(s)
}

func printSlice(tag string, slice []int)    {
    fmt.Printf("%s len=%d cap=%d %v\n", tag, len(slice), cap(slice), slice)
}

