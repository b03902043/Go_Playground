package pointer

import (
    "fmt"
    "time"
    "golang.org/x/tour/tree"
    _ "sync"
)

func say(message string)    {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(message)
    }
}

func sum(array []int, c chan int)  {
    _sum := 0
    for i := range array    {
        _sum += array[i]
    }
    c <- _sum // send result to channel c
}

func count(N int, c chan int)   {
    for i := 0; i < N; i++  {
        c <- i
    }
    close(c)
}

func count_with_quit(c chan int, quit chan int)  {
    x := 777
    for {
        select {
        case c <- x:
            x += 1
        case <- quit:
            return
        }
    }
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int)	{
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int)	{
	if t == nil	{
		return
	}

	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool	{
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		a, ok1 := <-ch1
		b, ok2 := <-ch2
		switch	{
			case ok1 != ok2:
				return false
			case ok1 == false:
				return true
			case a != b:
				return false
		}
	}
}

func RunRoutine()   {
    go say("hi")
    say("Interruptable")

    s := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c
    fmt.Printf("%d %d %d\n", x, y, x+y)

    buffered_channel := make(chan int, 2)
    buffered_channel <- 1
    buffered_channel <- 1
    x, y = <-buffered_channel, <-buffered_channel
    fmt.Printf("%d %d\n", x, y)

    ch := make(chan int, 5)
    go count(cap(ch), ch)
    for val := range ch {
        fmt.Println(val)
    }

    fmt.Println("Select demo")
    c = make(chan int)
    quit := make(chan int)
    go func()   {
        for i := 0; i < 10; i ++    {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    count_with_quit(c, quit)

    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(2), tree.New(1)))
}
