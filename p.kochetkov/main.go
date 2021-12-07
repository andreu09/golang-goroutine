package main

import (
	"fmt"
	"math/rand"
	"time"
)

var iterations int

func main() {

	fmt.Print("Введите размер канала: ")
	fmt.Scan(&iterations)

	ch1 := make(chan int, iterations)
	ch2 := make(chan bool)

	var n1, n2 int

	go producer(ch1)
	go consumer(ch1, &n2, &n1, ch2)

	<-ch2

	fmt.Println("Кол-во нечетных чисел: ", n1)
	fmt.Println("Кол-во четных чисел: ", n2)

}

func producer(ch chan int) {

	for i := 0; i < iterations; i++ {

		rand.Seed(time.Now().UnixNano())
		ch <- rand.Intn(100)
	}
}

func consumer(ch chan int, n1 *int, n2 *int, exit chan bool) {

	for i := 0; i < iterations; i++ {
		if <-ch%2 != 0 {
			*n1++
		} else {
			*n2++
		}

	}
	exit <- true

}
