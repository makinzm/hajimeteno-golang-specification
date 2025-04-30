package main

import (
  "fmt"
  "time"
)

// 送信専用チャネルを引数に取る関数
func producer(ch chan<- int) {
  for i := 0; i < 5; i++ {
    ch <- i  // チャネルに値を送信
    fmt.Printf("送信: %d\n", i)
  }
  close(ch)  // チャネルを閉じる
}

// 受信専用チャネルを引数に取る関数
func consumer(ch <-chan int) {
  // チャネルからの受信（チャネルが閉じられるまで）
  for {
    val, ok := <-ch
    if !ok {
      // チャネルが閉じられた場合
      fmt.Println("チャネルは閉じられました")
      return
    }
    fmt.Printf("受信: %d\n", val)
    time.Sleep(100 * time.Millisecond) // 受信側が遅いケース
  }
}

func main() {
  // バッファなしチャネルの例
  fmt.Println("=== バッファなしチャネル ===")
  unbuffered := make(chan int)

  // NOTE: 順番を変えると、mainプロセスがDeadLockになるため注意
  go producer(unbuffered)
  consumer(unbuffered)

  // バッファありチャネルの例
  fmt.Println("\n=== バッファありチャネル（容量2） ===")
  buffered := make(chan int, 2)
  
  // バッファ内にどれだけ値が格納されているか確認
  fmt.Printf("バッファ使用量: %d, 容量: %d\n", len(buffered), cap(buffered))
  
  // バッファが満杯になるまでブロックしない
  buffered <- 10
  fmt.Printf("10を送信後 - バッファ使用量: %d\n", len(buffered))
  
  buffered <- 20
  fmt.Printf("20を送信後 - バッファ使用量: %d\n", len(buffered))
  
  // バッファから値を取り出す
  val1 := <-buffered
  fmt.Printf("受信: %d, バッファ使用量: %d\n", val1, len(buffered))
  
  val2 := <-buffered
  fmt.Printf("受信: %d, バッファ使用量: %d\n", val2, len(buffered))
}
