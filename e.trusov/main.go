package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a1 := 0
	a2 := 0
	ch1 := make(chan int, 100000)
	done := make(chan bool)

	go produser(ch1)
	go consumer(&a1, &a2, ch1, done)
	<-done

	fmt.Println("yes")
	fmt.Println("Четных:", a1)
	fmt.Println("Нечетных", a2)
}
func produser(ch1 chan int) { //генерация чисел

	for i := 0; i < 100000; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)
		ch1 <- n
		//fmt.Println(n)
	}

}

func consumer(a1 *int, a2 *int, ch1 chan int, done chan bool) { // подсчет четных/нечетных

	for i := 0; i < 100000; i++ {

		if <-ch1%2 == 0 {
			*a1 = *a1 + 1
		} else {
			*a2 = *a2 + 1
		}

	}
	done <- true
}
