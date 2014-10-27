package main

import (
    "fmt"
  )

  func main() {
    chanel := make(chan int)

    go func(s chan<- int) {
      for i := 1; i < 5; i++ {
        s <- i
      }
      close(s)
    }(chanel)

    for {
      value, ok := <-chanel
      if !ok {
        fmt.Println("fin")
        break
      }
      fmt.Println(value)
    }
  }
