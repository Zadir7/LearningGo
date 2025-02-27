package main

import "fmt"

func main() {
	fmt.Println("hello, world")

	var a int = 7
	var b = 4
	c, d := 3, 15
	fmt.Printf("%d %d %d %d\n", a, b, c, d)

	const (
		e = iota
		f = iota
		g
		h
	)
	fmt.Printf("%d %d %d %d\n", e, f, g, h)

	arr := []int{3, 5, 7, 9}
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}

	if a > b {
		fmt.Printf("\nmore\n")
	} else {
		fmt.Printf("\nless\n")
	}

}
