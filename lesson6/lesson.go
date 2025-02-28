package main

import (
	"fmt"
	"sync"
)

type Age int //size = 8

type User struct {
	Id      int64   `json:"-"`        //Тег поля - игнорирование в encode/json. Получить информацию можно через reflect
	Name    string  `gorm:"size:255"` //GORM - ORM
	Age     Age     //size = 8, offset = 24
	friends []int64 //Приватные поля тоже можно достать через reflect
}

func (u User) IsOldEnough() bool { //Передача по значению
	return u.Age > 18
}

func (u *User) ChangeName(newName string) { //Передача по ссылке
	u.Name = newName
}

// Реализовать тип IntStack с методами Push(i int) и Pop() int
type IntStack struct {
	ints []int
}

func (stack *IntStack) Push(i int) {
	stack.ints = append(stack.ints, i)
}

func (stack *IntStack) Pop() int {
	stackLength := len(stack.ints)
	if stackLength == 0 {
		panic("Trying to pop empty stack")
	}
	lastItem := stack.ints[stackLength-1]
	stack.ints = stack.ints[0 : stackLength-1]
	return lastItem
}

type LinkStorage struct {
	sync.Mutex //Встроенный тип
	storage    map[string]string
}

// вместо
//storage.Mutex.Lock()
// можно просто
//storage.Lock()

//При этом, это не наследование, методы не переопределяются
/*
type Base struct {}
func (b Base) Name() string {
 return "Base"
}
func (b Base) Say() {
 fmt.Println(b.Name())
}
type Child struct {
 Base
 Name string
}
func (c Chind) Name() string {
 return "Child"
}
c := Child{}
c.Say() // Base
*/

func main() {

	//u1 := User{} - Zero Value
	//u2 := &User{} - Zero Value (pointer)
	//var wordCount []struct {word string; count int} - Anonymous type

	x := 1
	xPtr := &x //Pointer
	y := *xPtr //Dereferencing pointer/Разыменование указателя
	fmt.Printf("%d\n", y)

	//var n *int //nil
	//nPtr := *n //panic

	u3 := User{Name: "Alexander", Age: 30}
	u3Ptr := &u3
	fmt.Printf("%s of age %d\n", u3Ptr.Name, u3Ptr.Age) //Обращение к полям структуры без разыменования

	stack := IntStack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	stack.Push(4)
	fmt.Printf("%d", stack.ints)
}
