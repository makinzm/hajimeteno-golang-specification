package main

import (
  "fmt"
)

func submain() string {
  defer func() {
    fmt.Println("Deferred function executed 001")
  }()
  defer func() {
    fmt.Println("Deferred function executed 002")
  }()
  return "Hello, World!"
}

func main() {
  submain_result := submain()
  fmt.Println(submain_result)
}
