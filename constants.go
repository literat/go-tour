package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Ahoj", World)
	fmt.Println("Štastný", Pi, "den")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}