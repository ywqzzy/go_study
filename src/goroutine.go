package main

import (
	"fmt"
	"runtime"
	"time"
)

func test1() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			//for {
			//	fmt.Printf("Hello from " + "goroutine %d\n", i)
			//}
			for {
				a[i]++
				runtime.Gosched() // 手动交出控制权
			}
		}(i) // 这样是不会打出东西的
	}

	// data race
	time.Sleep(time.Millisecond) // 这样才能打出来...
	fmt.Println(a)
	// 协程  轻量级线程
	// 非抢占式多任务处理， 由协程主动交出控制权
	// 编译器/解释器/虚拟机层面的多任务
	// 多个协程可以在一个或多个线程上运行

}

func main() {
	// 子程序是协程的一个特例c
}
