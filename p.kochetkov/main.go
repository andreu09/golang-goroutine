package main

import (
	"fmt"
	"math/rand"
	"time"
)

var iterations int = 100000

func main() {

	first := make(chan int, 100000)
	second := make(chan bool)

	var evenCount int
	var oddCount int

	go producer(first)
	go consumer(first, &evenCount, &oddCount, second)

	<-second

	fmt.Println("четных: ", evenCount)
	fmt.Println("нечетных: ", oddCount)

}

func producer(ch chan int) {
	for i := 0; i < iterations; i++ {
		rand.Seed(time.Now().UnixNano())
		ch <- rand.Intn(30)
	}

}

func consumer(ch chan int, eC *int, oC *int, exit chan bool) {
	for i := 0; i < iterations; i++ {
		if <-ch%2 == 0 {
			*eC++
		} else {
			*oC++
		}
	}
	exit <- true
}
