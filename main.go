package main

import "fmt"

// 隠れた依存関係のある変数
type I interface { ab() []int }
type T struct{}

func (T) ab() []int { return []int{hidden_a, hidden_b} }

var hidden_x = I(T{}).ab()  // hidden_xには、hidden_aとhidden_bへの検出されない隠れた依存関係がある
// var _ = sideEffect()
var hidden_a = hidden_b
var hidden_b = 42

// func sideEffect() bool {
//   fmt.Println("sideEffect()が呼び出されました")
//   return true
// }

func main() {
  fmt.Println("\n=== 隠れた依存関係 ===")
  fmt.Printf("hidden_x = %v\n", hidden_x)
  fmt.Printf("hidden_a = %d\n", hidden_a)
  fmt.Printf("hidden_b = %d\n", hidden_b)
  fmt.Println("注: hidden_xの初期化タイミングは不定")
}

