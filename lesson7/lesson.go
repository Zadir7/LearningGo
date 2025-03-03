package main

import (
	"fmt"
	"strconv"
)

//Интерфейс это набор методов. Если тип реализует все указанные методы, значит он удовлетворяет интерфейсу.
//Интерфейсы реализуются неявно.
//Один тип может реализовать много интерфейсов, один интерфейс может релизоваться множеством типов.
//Интерфейс может встраивать другой интерфейс
/*
type Greeter interface {
     hello()
}
type Stranger interface {
    Bye() string
    Greeter
    fmt.Stringer
}
*/
//Имена методов не должны повторяться (нет перегрузок)
type Empty interface{} //пустой интерфейс, ему удовлетворяет любой тип.

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

type I1 interface{ M1() }
type T1 struct{}

func (T1) M1() {}

type I2 interface {
	I1
	M2()
}
type T2 struct{}

func (T2) M1() {}
func (T2) M2() {}

type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}

func main() {
	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)

	var x fmt.Stringer
	//Значение типа интерфейс, состоит из динамического типа и значения, нулевое значение для обоих - nil
	//Такое значение nil только когда ОБЕ его части nil
	x = Temp(24)
	fmt.Printf("%v %T\n", x, x) // 24 °C main.Temp

	//Type assertion
	//x.(T) проверяет, что конкретная часть значения x имеет тип T и x != nil
	//­если T ­ не интерфейс, то проверяем, что динамический тип x это T
	//если T ­ интерфейс: то проверяем, что динамический тип x его реализует
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)      // hello
	s, ok := i.(string) // hello true
	fmt.Println(s, ok)
	r, ok := i.(fmt.Stringer) // <nil> false
	fmt.Println(r, ok)
	f, ok := i.(float64) // 0 false
	fmt.Println(f, ok)

	//f = i.(float64) // panic: interface conversion

	//Проверка нескольких типов через switch
	var v I1
	switch v.(type) {
	case T1:
		fmt.Println("T1")
	case T2:
		fmt.Println("T2")
	case nil:
		fmt.Println("nil")
	default:
		fmt.Println("default")
	}

}
