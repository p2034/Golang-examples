package main

import "fmt"

// https://www.callicoder.com/golang-slices/

func example_modify() {
	// slice
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // [2 3 5 7 11 13]

	// add new element
	primes = append(primes, 17)
	fmt.Println(primes) // [2 3 5 7 11 13 17]

	// change
	primes[3] = 29
	fmt.Println(primes) // [2 3 5 29 11 13 17]

	// add slice
	notprimes := []int{1, 4, 6, 8, 9}
	numbers := append(primes, notprimes...)
	fmt.Println(numbers) // [2 3 5 29 11 13 17 1 4 6 8 9]
}

func example_from_array() {
	// array
	primes_arr := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes_arr)

	// slice
	primes := primes_arr[0:3]
	fmt.Println(primes, " - [0:3]")

	primes = primes_arr[:5]
	fmt.Println(primes, " - [:5]")

	primes = primes_arr[3:]
	fmt.Println(primes, " - [3:]")

	primes = primes_arr[:]
	fmt.Println(primes, " - [:]")
}

func example_from_make() {
	// type, length, capacity
	str := make([]string, 2, 4)
	fmt.Println(str)
	fmt.Println(len(str), cap(str))
}

func example_copy() {
	src := []string{"C", "C++", "Golang", "Swift", "Java", "JavaScript", "PHP", "Python"}
	var dest = make([]string, 4)

	num := copy(dest, src)
	fmt.Println(num, dest) // 4 [C C++ Golang Swift]
}

func example_slice() {
	sss := [][]string{
		{"China, Russia", "Japan", "India"},
		{"Brasil", "USA", "Canada", "Argentina"},
	}
	fmt.Println(sss)
}

func main() {
	example_slice()
}
