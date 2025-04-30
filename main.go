package main

import "fmt"

// 基本構造体
type Base struct {
  Name string
}

// Baseのメソッド
func (b Base) GetName() string {
  return b.Name
}

func (b *Base) SetName(name string) {
  b.Name = name
}

// 埋め込みを使った構造体
type Derived struct {
  Base  // 埋め込み（深さ1）
  Age int
}

// さらに埋め込みを使った構造体
type SuperDerived struct {
  Derived  // 埋め込み（Base は深さ2）
  Job string
}

// 同じ名前のフィールドを持つ構造体
type Conflicting struct {
  Base1 Base // 名前の競合
  Base2 Base // 名前の競合
}

// インターフェース
type Named interface {
  GetName() string
}

func main() {
  // 1. 基本的なセレクタの使用
  base := Base{Name: "John"}
  fmt.Println(base.Name)       // フィールドへのアクセス
  fmt.Println(base.GetName())  // メソッドへのアクセス

  // 2. ポインタ経由のセレクタ
  basePtr := &Base{}
  fmt.Println(basePtr.Name)      // (*basePtr).Name の省略形
  basePtr.Name = "Doe"
  fmt.Println(basePtr.GetName()) // 値レシーバーのメソッドもポインタから呼び出せる

  // 3. 埋め込みとセレクタ
  derived := Derived{
    Base: Base{Name: "Bob"},
    Age:  30,
  }
  fmt.Println(derived.Name)      // 埋め込まれたBaseのフィールドにアクセス
  fmt.Println(derived.GetName()) // 埋め込まれたBaseのメソッドにアクセス

  // 4. 多段の埋め込みとセレクタ
  superDerived := SuperDerived{
    Derived: derived,
    Job:     "Engineer",
  }
  fmt.Println(superDerived.Name)      // 2段階の埋め込みを通してアクセス
  fmt.Println(superDerived.GetName()) // 2段階の埋め込みを通してメソッドにアクセス
  fmt.Println(superDerived.Age)       // 直接埋め込まれたDerivedのフィールドにアクセス

  // 5. 名前の競合 -> Promotionの失敗
  // conflictingFailed := Conflicting{
  //   Base1: Base{Name: "Original"},
  //   Base2: Base{Name: "Conflict"},
  // }
  // fmt.Println(conflictingFailed.Name) // Compile Error: conflictingFailed.Name undefined
  
  // 明示的に指定する必要がある
  conflicting := Conflicting{
    Base1: Base{Name: "Original"},
    Base2: Base{Name: "Conflict"},
  }
  fmt.Println(conflicting.Base1.Name)  // Base.Name を明示
  fmt.Println(conflicting.Base2.Name) // Other.Name を明示

  // 6. nilポインタとセレクタ（実行時パニック）
  // var nilPtr *Base = nil
  // fmt.Println(nilPtr.Name) // パニック: ランタイムエラー

  // 7. nilインターフェースとセレクタ（実行時パニック）
  // var nilInterface Named = nil
  // fmt.Println(nilInterface.GetName()) // パニック: ランタイムエラー
}
