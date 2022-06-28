package main

import "fmt"

// I just want to show here named return

func example_func() (name string, age int) {
	age = 34
	name = "Jason"
	return
}

func main() {
	n, a := example_func()
	fmt.Println(n, a)
}
