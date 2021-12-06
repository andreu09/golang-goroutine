package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		chet   int
		nechet int
	)
	c1 := make(chan int, 100000)
	c2 := make(chan bool)
	go producer(c1)
	go consumer(c1, c2, &chet, &nechet)

	<-c2

}

func producer(canal chan int) {
	rand.Seed(time.Now().UnixNano())
	var a int
	for i := 0; i < 100000; i += 1 {
		a = rand.Int()
		fmt.Print(a)
		canal <- a
	}

}

func consumer(canal chan int, c2 chan bool, chet *int, nechet *int) {
	for b := 0; b < 100000; b += 1 {
		c := <-canal
		if c%2 == 0 {
			*chet += 1
		} else {
			*nechet += 1
		}
		fmt.Print("Consumer work")
	}
	fmt.Print("Work done \n")
	c2 <- true
}
