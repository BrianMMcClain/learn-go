package main

import (
  "fmt"
  "time"
)

func ping(c chan string) {
  for {
    c <- "ping"
    time.Sleep(10000)
  }
}

func pong(c chan string) {
  for {
    c <- "pong"
    time.Sleep(10000)
  }
}

func main() {
  fmt.Println("Building channel")
  c := make(chan string, 100)

  go ping(c)
  go pong(c)
  s := ""

  for {
    select {
    case s = <-c:
      fmt.Println(s)
    default:
    }
  }
}