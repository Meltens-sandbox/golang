// Package main provides ...
package main

//goroutine非同期処理

import (
	"log"
	"time"
)

func main() {
	c := make(chan int)
	d := make(chan struct{}, 1)

	go func() {
		defer close(d)
		for i := 1; i <= 3; i++ {
			c <- i
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for v := range c {
			log.Println(v)
		}
	}()
	<-d
}
