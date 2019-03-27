package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(10)
	fmt.Println("Moje oblíbené číslo je", rand.Intn(10))
}