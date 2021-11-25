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
	syncing := make(chan bool)
	// Четных элементы
	evenNum := 0
	// Не четных элементы
	noEvenNum := 0

	go producer(transmission)
	go consumer(&evenNum, &noEvenNum, syncing, transmission)
	// Чтобы программа не завершилась до запуска горутин
	<-syncing
	fmt.Print("Четных: ", evenNum, "\nНе четных: ", noEvenNum, "\n")

}

func producer(transmission chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < dimensionTransmissionClannel; i++ {
		transmission <- rand.Intn(10000)
	}
}

func consumer(evenNum *int, noEvenNum *int, syncing chan bool, transmission chan int) {
	for i := 0; i < dimensionTransmissionClannel; i++ {
		if <-transmission%2 == 0 {
			*evenNum += 1
		} else {
			*noEvenNum += 1
		}
	}
	// Чтобы избежать deadlock говорим каналу синхронизации что мы закончили операцию
	syncing <- true
}
