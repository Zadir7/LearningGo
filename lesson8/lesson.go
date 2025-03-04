package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var ch = make(chan struct{})

	go func() {
		fmt.Printf("Hello")
		close(ch) //Синхронизация
	}()

	//f(out chan<- int) канал только для записи
	//f(in <-chan int) канал только для чтения

	for {
		x, ok := <-ch
		fmt.Println(x)
		if !ok {
			break
		}
	} //Читаем канал пока не закрыт
	//for x := range ch

	/*
	   Select ждет когда один из каналов будет открыт на чтение или запись, если никто не готов, сработает default
	   Пустой select{} ждет вечно
	       select {
	       case <-ch1:
	       // ...
	       case ch2 <- y:
	       // ...
	       default:
	       // ....
	       }
	*/

	//timer := time.NewTimer(10 * time.Second) Таймер: срабатывает 1 раз через указанное время
	//ticker := time.NewTicker(10*time.Second) Тикер: срабатывает раз в заданное количество времени

	//Graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	fmt.Printf("Got %v...\n", <-interrupt)
	//Выключаем веб-сервер и обработчики задач

}
