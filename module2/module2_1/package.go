package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(3)
	fmt.Println("My favorite number is", rand.Intn(10))
}
