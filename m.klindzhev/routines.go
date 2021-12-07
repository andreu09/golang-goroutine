package main

import (
	"fmt"
	"math/rand"
)

func main() {
	iterations := 100000
	first_channel := make(chan int, iterations)
	second_channel := make(chan bool)

	even := 0
	odd := 0

	go producer(first_channel, iterations)
	go consumer(first_channel, second_channel, &even, &odd, iterations)

	<-second_channel

	fmt.Println("Even:", even)
	fmt.Println("Odd:", odd)
}

func producer(first_channel chan int, iterations int) {
	for i := 0; i < iterations; i++ {
		first_channel <- rand.Intn(99)
	}
}

func consumer(first_channel chan int, second_channel chan bool, even *int, odd *int, iterations int) {
	for i := 0; i < iterations; i++ {
		if <-first_channel%2 == 0 {
			*even++
		} else {
			*odd++
		}
	}
	second_channel <- true
}
