package main

import "fmt"
import "rsc.io/quote"

func greet(name string) string {
  msg := fmt.Sprintf("Good day to you, %v!", name)
  return msg
}

func main() {
  greet("Steven")
  fmt.Println(quote.Glass())
}
