package main

import "math/rand"

func calculate() int {
	return rand.Intn(5) + rand.Intn(4)
}
