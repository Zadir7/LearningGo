package main

import (
	"fmt"
	"sync"
	"time"
)

type Dog struct {
	name         string
	walkDuration time.Duration
}

func (d Dog) Walk(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is taking a walk\n", d.name)
	time.Sleep(d.walkDuration)
	fmt.Printf("%s is going home\n", d.name)
}

func walkTheDogs(dogs []Dog) {
	var wg sync.WaitGroup

	for _, d := range dogs {
		wg.Add(1)
		d.Walk(&wg)
	}

	wg.Wait()
	fmt.Println("everybody's home")
}

var i int // i == 0
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	i = i + 1
	m.Unlock() // release lock
	wg.Done()
}

func (d *Dog) Bark() { fmt.Printf("%s", d.name) }

var dogPack = sync.Pool{
	New: func() interface{} { return &Dog{} },
}

type OldDog struct {
	name string
	die  sync.Once
}

func (d *OldDog) Die() {
	d.die.Do(func() { println("bye!") })
}

func main() {
	//Sync.WaitGroup (looks similar to Tasks.WaitAll in .net)
	/*
		dogs := []Dog{{"vasya", time.Second}, {"john", time.Second * 3}}
		walkTheDogs(dogs)
	*/

	//Mutex
	/*
		var wg sync.WaitGroup
		var m sync.Mutex
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go worker(&wg, &m)
		}
		wg.Wait()
		fmt.Println("value of i after 1000 operations is", i)
	*/

	//sync.Pool
	//Temp storage, goroutine-safe
	//No guarantee of object lifetime + loads GC
	//Lowers time for memory allocation
	/*
		dog := dogPack.Get().(*Dog)
		dog.name = "ivan"
		dog.Bark()
		dogPack.Put(dog)
	*/

	//sync.Once
	//Guarantees only 1 execution
	/*
		d := OldDog{name: "bob"}
		d.Die()
		d.Die()
	*/
}

// RWMutex
type Counters struct {
	mx sync.RWMutex
	m  map[string]int
}

func (c *Counters) Load(key string) (int, bool) {
	c.mx.RLock() //Locks only reading
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

//Also there is sync.Map if system is so high-load and has 32+ cores
//that 100ns is critical
