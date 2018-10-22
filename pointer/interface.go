package pointer

import (
    "fmt"
    "time"
)

/*  fmt.Stringer:
    type Stringer interface {
        String() string
    }
*/

type Adder interface    {
    Add(b int)
}

type Spin struct    {
    v int
}

func (s Spin) String() string   {
    return fmt.Sprintf("%s(%d)", "Spin", s.v)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// type = error
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func (s Spin) Add(a int)   {
    s.v += a
}

func InterfaceDemo()    {
    var s Adder
    s = Spin{1}
    s.Add(10)
    describe(s)

    var I interface{}
    I = "1"
    describe(I)

    // type assertion
    value, ok := I.(Spin)
    fmt.Println(value, ok)

    // type switch
    switch v := I.(type) {
    case int:
        fmt.Println(v)
    default:
        fmt.Println("default", v)
    }

    if err := run(); err != nil {
        fmt.Println(err)
    }
}

func describe(i interface{})    {
    fmt.Printf("%v %T\n", i, i)
}
