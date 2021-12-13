package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	id             int
	leftChopstick  *chopstick
	rightChopstick *chopstick
}

func (p philosopher) eat(wg *sync.WaitGroup, currentEatersCountChannel chan int) {
	fmt.Println("Philosopher", p.id, "eat()")
	for i := 0; i < 3; i++ {
		currentEatersCountChannel <- p.id

		// randomize picking chopsticks
		randBinaryValue := rand.Intn(2)
		// fmt.Println(randBinaryValue)
		if randBinaryValue == 0 {
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()
		} else {
			p.rightChopstick.Lock()
			p.leftChopstick.Lock()
		}

		fmt.Println("Philosopher", p.id, ": Starting to eat - round", i+1)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Philosopher", p.id, ": Finishing eating - round", i+1)

		if randBinaryValue == 0 {
			p.rightChopstick.Unlock()
			p.leftChopstick.Unlock()
		} else {
			p.leftChopstick.Unlock()
			p.rightChopstick.Unlock()
		}

		<-currentEatersCountChannel
	}
	fmt.Println("Philosopher", p.id, "~~~~~~~~~eat()")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())
	currentEatersCountChannel := make(chan int, 2)

	chopsticks := make([]*chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(chopstick)
	}

	wg.Add(5)

	philosophers := make([]*philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &philosopher{
			i + 1,
			chopsticks[i],
			chopsticks[(i+1)%5],
		}
		go philosophers[i].eat(&wg, currentEatersCountChannel)
	}

	wg.Wait()
}
