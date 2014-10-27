package main

import (
    "fmt"
)

func main() {

  a := make(chan int)
  b := make(chan int)
  c := make(chan int)

  go func() {
    for {
      fmt.Print(".")
      a <- 0
    }
  }()

  go func() {
    for {
      fmt.Print("*")
      b <- 0
    }
  }()

  go func() {
    for {
      fmt.Print("_")
      c <- 0
    }
  }()

  for i := 0; i < 10; i++ {
    select {
      case <-a:
        fmt.Println("catch A")
      case <-b:
        fmt.Println("catch B")
      case <-c:
        fmt.Println("catch C")
    }
  }
}
