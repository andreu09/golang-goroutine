package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Размерность буферизированного канала
var dimensionTransmissionClannel int = 10000

func main() {
	// Буферизированный канал передачи данных
	transmission := make(chan int, dimensionTransmissionClannel)
	// Канал синхронизации
	syncing1 := make(chan bool)
	syncing2 := make(chan bool)
	// Четных элементы
	evenNum := 0
	// Не четных элементы
	noEvenNum := 0

	go producer1(transmission)
	go producer2(transmission)
	go consumer1(&evenNum, &noEvenNum, syncing1, transmission)
	go consumer2(&evenNum, &noEvenNum, syncing2, transmission)
	<-syncing1
	<-syncing2
	// Чтобы программа не завершилась до запуска горутин

	fmt.Print("Четных: ", evenNum, "\nНе четных: ", noEvenNum, "\n")

}

func producer1(transmission chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < dimensionTransmissionClannel; i++ {
		transmission <- rand.Intn(10000)
	}
}

func producer2(transmission chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < dimensionTransmissionClannel; i++ {
		transmission <- rand.Intn(10000)
	}

}

func consumer1(evenNum *int, noEvenNum *int, syncing1 chan bool, transmission chan int) {
	for i := 0; i < dimensionTransmissionClannel; i++ {
		if <-transmission%2 == 0 {
			*evenNum += 1
		} else {
			*noEvenNum += 1
		}
	}
	// Чтобы избежать deadlock говорим каналу синхронизации что мы закончили операцию
	syncing1 <- true
}

func consumer2(evenNum *int, noEvenNum *int, syncing2 chan bool, transmission chan int) {
	for i := 0; i < dimensionTransmissionClannel; i++ {
		if <-transmission%2 == 0 {
			*evenNum += 1
		} else {
			*noEvenNum += 1
		}
	}
	// Чтобы избежать deadlock говорим каналу синхронизации что мы закончили операцию
	syncing2 <- true
}
