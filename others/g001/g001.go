// 交替打印奇偶数
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	notify1, notify2 := make(chan bool, 1), make(chan bool)
	notify1 <- true
	go func() {
		for i := 1; i <= 10; i += 2 {
			<-notify1
			fmt.Println(i)
			notify2 <- true
		}
		wg.Done()
	}()

	go func() {
		for i := 2; i <= 10; i += 2 {
			<-notify2
			fmt.Println(i)
			notify1 <- true
		}
		wg.Done()
	}()

	wg.Wait()
}

