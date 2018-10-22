package pointer

import (
    "fmt"
    "math"
)

type float float64

type Vertex struct  {
    x, y float64
}

// method can only on local types
func (f *float) sqrt() float64 {
    return math.Sqrt(float64(*f))
}

// pass by value
func Scale(v Vertex, rate float64)   {
    v.x *= rate
    v.y *= rate
}

// pass by reference
func (v *Vertex) Scale(rate float64)    {
    v.x *= rate
    v.y *= rate
}

func Method_usage() {
    val := float(16)
    fmt.Println(val.sqrt())
    v := Vertex{1, 2}
    //Scale(v, 10)
    v.Scale(10)
    fmt.Println(v)

}
