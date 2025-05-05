package timerz

import (
	"fmt"
	"sync"
	"time"
)

func Timer() {

	var wg sync.WaitGroup

	timer := time.NewTimer(2 * time.Second)
	<- timer.C	// when the timer is up, it will send message to channel C
	fmt.Println("helo")

	timer2 := time.NewTimer(2 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	wg.Wait()
}
