package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	//var arr [256]int - Array of 256 ints
	arr := []int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	slc1 := arr2[1:2]

	slc1Copy := make([]int, len(arr))
	copy(slc1Copy, arr)
	slc1Copy[0] = 99

	sort.Ints(slc1Copy)

	//arr[3] = 4 index out of range runtime error
	//arr2[3] = 4 compilation error
	fmt.Printf("%d\n", arr)
	fmt.Printf("%d\n", slc1)
	fmt.Printf("%d\n", slc1Copy)

	value, ok := words["go"]
	if !ok {
		fmt.Printf("%d\n", value)
	}
	//for key, val := range cache  iterates key-value pairs
	//for _, val := range cache  excludes keys
	//for key, _ := range cache  excludes values

	var seq []string                                   //initialized with nil
	seq = append(seq, "hello there", "general kenobi") //works fine
	fmt.Printf("%s\n", seq)
}

// Concurrency in maps
var words = map[string]int{"hello": 1, "world": 2}
var mutex sync.RWMutex

func Get(key string) int {
	var result int
	mutex.RLock()
	result = words[key]
	mutex.RUnlock()
	return result
}

func Set(key string, value int) {
	mutex.RLock()
	words[key] = value
	mutex.RUnlock()
}
