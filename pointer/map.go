package pointer

import (
    "fmt"
    "strings"
)

func Map_usage()    {

    map_from_make := make(map[string]int)
    map_from_make["hello"] = 100

    // map literal
    map_from_init := map[string]int {
        "a": 1,
        "b": 2,
    }
    delete(map_from_init, "a")
    elem, ok := map_from_init["a"]
    fmt.Printf("key \"%s\" exist? %v val? %v\n", "a", ok, elem)

    fmt.Println(map_from_make, map_from_init)
}

func WordCount(s string) map[string]int {
    dict := make(map[string]int)
    for _, v := range strings.Fields(s) {   // strings.Fields(s string)  ==  str.split() in python
        elem, _ := dict[v]
        dict[v] = elem + 1
    }
    return dict
}
