package main

import "fmt"

type Node struct {
    Value int
    Next  *Node
}

func main() {
    n1 := &Node{Value: 1}
    n2 := &Node{Value: 2}
    n3 := &Node{Value: 3}

    n1.Next = n2
    n2.Next = n3

    fmt.Println(n1.Value) // 1
    fmt.Println(n1.Next.Value) // 2
    fmt.Println(n1.Next.Next.Value) // 3

    selfLoop := &Node{Value: 1}
    selfLoop.Next = selfLoop

    fmt.Println(selfLoop.Value) // 1
    fmt.Println(selfLoop.Next.Value) // 1
    fmt.Println(selfLoop.Next.Next.Value) // 1
    fmt.Println(selfLoop.Next.Next.Next.Value) // 1
}

