package main

type A int
type B int

func calculateGeneral[T ~int](a T) A {
  return A(a * 2)
}

type Person struct {
  Name string
  Age  int
}

type Teacher struct {
  Name string
  Age  int
}


func main() {
  var a A = 5
  resultA := calculateGeneral(a)
  println(resultA) // Output: 10

  var b B = 10
  // a = b // Compile error: cannot use b (variable of type B) as A value in assignment
  resultB := calculateGeneral(b)
  println(resultB) // Output: 20

  teacher := Teacher{Name: "John", Age: 30}
  // var person Person = teacher // Compile error: cannot use teacher (variable of type Teacher) as Person value in assignment: Teacher does not implement Person
  person := Person(teacher) // Implicit conversion from Teacher to Person

  println(person.Name) // Output: John
  println(person.Age)  // Output: 30
}

