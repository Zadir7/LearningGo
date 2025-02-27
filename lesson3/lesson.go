package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello"
	s2 := s + " go"
	var b strings.Builder
	for _, r := range s2 {
		b.WriteRune(r)
	}
	fmt.Printf("%s", b.String())
}

type Person struct {
	Name string //public
	age  int    //private
}
