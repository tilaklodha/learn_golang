package main

import "fmt"

func main() {
	log("golang")
	add(2, 3)
	power()
}

//No return type
func log(message string) {
	fmt.Println("From log")
}

//One return value
func add(a int, b int) int {
	fmt.Println("From add")
	return a + b
}

//Two return value
func power() (int, bool) {
	fmt.Println("From power")
	return 1, true
}
