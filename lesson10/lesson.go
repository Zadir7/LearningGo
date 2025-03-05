package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := "./data.txt"
	file, _ := os.Create(path)
	b := make([]byte, 1024*1024)

	err := os.WriteFile(path, b, os.ModeDevice)
	if err != nil {
		log.Panicf("write failed: %v", err)
	}

	//read, _ := os.ReadFile(path)
	//fmt.Printf("%v", read) //1Mb of zeroes

	file.Close()

	cfs := CustomFormattedString{value: "null"}
	fmt.Printf("%v ", cfs)
	fmt.Printf("%#v ", cfs)
}

//fmt formatting
//%v default value
//%#v default value in Go code
//%T type
//%% 'escaped' % symbol
//%b numbers (binary)
//%o numbers (octet)
//%d numbers (decimal)
//%x numbers (hexadecimal (lower a-f))
//%s string
//%p memory address

type CustomFormattedString struct {
	value string
}

func (cfs CustomFormattedString) String() string {
	return "hello"
}

func (cfs CustomFormattedString) GoString() string {
	return "code"
}
