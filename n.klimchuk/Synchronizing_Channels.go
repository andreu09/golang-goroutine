package main

import (
	"fmt"
	"math/rand"
	"time"
)

var buffered_channel int = 100000

func main() {
	first_channel := make(chan int, buffered_channel)
	second_channel := make(chan bool)

	even_element := 0
	odd_element := 0

	go producer(first_channel)
	go consumer(first_channel, second_channel, &even_element, &odd_element)

	<-second_channel

	fmt.Println("Number of even elements =", even_element)
	fmt.Println()
	fmt.Println("Number of odd elements =", odd_element)
}

func producer(first_channel chan int) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < buffered_channel; i++ {
		first_channel <- rand.Intn(10)
	}
}

func consumer(first_channel chan int, second_channel chan bool, even_element *int, odd_element *int) {
	for i := 0; i < buffered_channel; i++ {
		if <-first_channel%2 == 0 {
			*even_element++
		} else {
			*odd_element++
		}
	}
	second_channel <- true
}
