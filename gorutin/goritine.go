package main

import (
	"fmt"
	"math/rand"
)

func main() {
	koll1 := 0
	koll2 := 0
	ch1 := make(chan int, 100000)
	ch2 := make(chan bool)

	go producer(ch1)
	go consumer(ch1, ch2, &koll1, &koll2)

	<-ch2

	fmt.Println("Чётных чисел: ", koll1)
	fmt.Println("Нечётных чисел: ", koll2)

}

func producer(ch1 chan int) {
	for i := 0; i < 100000; i++ {
		ch1 <- rand.Intn(101)
	}
}

func consumer(ch1 chan int, ch2 chan bool, koll1 *int, koll2 *int) {

	for k := 0; k < 100000; k++ {
		if <-ch1%2 == 0 {
			*koll1++
		} else {
			*koll2++
		}
	}
	ch2 <- true
}
