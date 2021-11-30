package main
import (
	"fmt"
	"time"
	"math/rand"
)

func producer(ch1 chan int){

	rand.Seed(time.Now().Unix())
	for i := 0; i < 100000; i++ {
		rn := rand.Intn(100)
		ch1 <- rn
	}
}

func consumer(ch1 chan int, ch2 chan bool, ev *int, od *int){

	for i :=0; i < 100000; i++{
		x :=<- ch1		
		if x %2 == 0{	
			*ev ++
		}else{
			*od ++
		}
	}		
	ch2 <- true
}

func main(){

	ch1 := make(chan int, 100000)
	ch2 := make(chan bool, 1)
	ev := 0
	od := 0
	
	go producer(ch1)
	go consumer(ch1, ch2, &ev, &od)

	<- ch2

	fmt.Println("even:", ev, "odd:", od)
}
