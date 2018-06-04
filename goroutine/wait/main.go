// Package main provides ...
package main

//goroutineにおいて処理が完了(done)するまで待機(wait)させる

import (
	"log"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		defer close(done)
		for i := 0; i < 3; i++ {
			log.Println("DONE")
			time.Sleep(1 * time.Second)
		}
	}()
	<-done

	log.Println("finish")
}
