package main

import (
  "fmt"
)

func main() {
  // Arrange
  one := 1
  zero_point_one := 0.1
  // Check type
  fmt.Printf("Type of one: %T\n", one)
  fmt.Printf("Type of zero_point_one: %T\n", zero_point_one)

  min_value := min(zero_point_one, one) // Failed : invalid argument: mismatched types float64 (previous argument) and int (type of one)
  max_value := max(zero_point_one, one) // Failed : invalid argument: mismatched types float64 (previous argument) and int (type of one)
  // min_value := min(1, 0.1)
  // max_value := max(1, 0.1)
  fmt.Println("Min:", min_value)
  fmt.Println("Max:", max_value)
  // Check type
  fmt.Printf("Type of min_value: %T\n", min_value)
  fmt.Printf("Type of max_value: %T\n", max_value)
}

