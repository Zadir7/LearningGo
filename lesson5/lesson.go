package main

import (
	"fmt"
	"net/http"
)

var packageBlockVar string

func NamedFunc() string {
	packageBlockVar = ""
	return packageBlockVar
}

func MultipleReturns(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func VariadicFunc(strs ...string) {
	for _, str := range strs {
		fmt.Printf("%s", str)
	}
}

func Closure() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

type Age int

func (age *Age) increaseAge() {
	*age++
}

func a() {
	i := 0
	defer fmt.Println(i) //0
	i++
	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
}

func main() {
	var a Age
	a = 18
	a.increaseAge()

	//defer fmt.Println(a) runs at the end of function
	//multiple defers execute in LIFO order

	//Practical closures
	http.HandleFunc("/index", hello)
	http.ListenAndServe(":3000", nil)

	/*v := "outer"
	fmt.Println(v) //outer
	{
		v := "inner"
		fmt.Println(v) //inner
		{
			fmt.Println(v) //inner
		}
	}
	{
		fmt.Println(v) //outer
	}
	fmt.Println(v) //outer

	anonFunc := func() {
		fmt.Println("Hello there")
	}
	anonFunc()

	seq1 := Closure()
	fmt.Println(seq1())
	fmt.Println(seq1())
	fmt.Println(seq1())

	seq2 := Closure()
	fmt.Println(seq2())*/
}
