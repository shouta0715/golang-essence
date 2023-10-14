package main

import (
	"fmt"
	"time"
)

func main() {

	ch := fanIn(generator("Hello"), generator("Bye"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

}

func fanIn(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				new_ch <- s
			case s := <-ch2:
				new_ch <- s
			default:
				fmt.Println("No value ready, moving on.")
			}
		}
	}()

	return new_ch
}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(1 * time.Second)
		}
	}()

	return ch
}
