package main

import (
	"fmt"
	"math/rand"
)

func main() {
  fmt.Println("Hello! This is a test")
  fmt.Println("Here's a random number: ", calculate())
}

func calculate() int {
  return rand.Intn(5) + rand.Intn(4)
}
