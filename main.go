package main

import (
  "fmt"
)

type Person struct {
  name string
  age  int
}

func main() {
  new_int := new(int)
  *new_int = 42
  fmt.Println(*new_int)
  normal_int := 42
  fmt.Println(&normal_int == new_int)

  new_person := new(Person)
  new_person.name = "John"
  new_person.age = 30  
  normal_person := Person{
    name: "John",
    age:  30,
  }
  fmt.Println(new_person == &normal_person)
}

