package main

import "fmt"

func main() {
	// For klasik
	for i := 1; i <= 5; i++ {
		fmt.Println("Ke:", i)
	}

	// Range
	angka := []int{10, 20, 40, 60, 80}
	for i, v := range angka {
		fmt.Println(i, "=", v)
	}
}
