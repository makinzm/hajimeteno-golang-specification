package main

import "fmt"

type T struct{}
func (T) Foo() {
  fmt.Println("T.Foo()")
}
func (*T) Bar() {
  fmt.Println("*T.Bar()")
}

type S struct {
  T
}

type SDash struct {
  *T
}

func main() {
    var s S
    s.Bar()
    s.Foo()

    var sdash SDash
    sdash.Bar()
    sdash.Foo()
}

