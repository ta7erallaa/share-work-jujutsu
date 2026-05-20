package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func divide(x, y int) int {
	return x / y
}

func multiple(x, y int) int {
	fmt.Println("I'm multiple function")
	return x * y
}
