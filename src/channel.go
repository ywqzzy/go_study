package main

import (
	"fmt"
	"time"
)

func chanDemo01() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
} // maybe cant print 1 and 2

func worker(id int, c chan int) {
	for {
		n := <-c
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func chanDemo02() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
} // maybe cant print 1 and 2

func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	return c
	// chan<- 只能发数据进channel
	// <-chan 只能收数据
}

func chanDemo03() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
} // maybe cant print 1 and 2

func bufferedChannel() {
	c := make(chan int, 3)
	go work03(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func createWorker01(id int) chan<- int {
	c := make(chan int)
	go work03(id, c)
	return c
}

func work03(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %c\n", id, <-c)
	}
}

func channleClose() {
	c := make(chan int)
	//go work03(0, c)
	//go work04(0,c)
	go work05(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func work04(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func work05(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func main() {
	//bufferedChannel()\
	channleClose()
}

// 不要通过共享内存来通信， 通过通信来共享内存
