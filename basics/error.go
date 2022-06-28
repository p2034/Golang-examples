package main

import "fmt"

func example_function() (string, error) {
	return "some text", nil
}

func example_defer() {
	// like in stack 4 3 2 1
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("...")
}

func example_panic_helper(depth int) {
	if depth > 0 {
		defer fmt.Println("In defer:", depth)
		fmt.Println("Before:", depth)
		example_panic_helper(depth - 1)
		fmt.Println("After:", depth)
	} else {
		defer fmt.Println("In defer from panic:", depth)
		panic("AAAAA")
	}
}

func example_panic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover:", r)
		}
	}()
	fmt.Println("Start recursion:")
	example_panic_helper(2)
	fmt.Println("End recursion.")
}

func main() {
	example_panic()
	fmt.Println("All is ok!")
}
