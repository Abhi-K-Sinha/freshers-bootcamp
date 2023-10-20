package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func addRating(chnl chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	delay := rand.Intn(100) + 1
	time.Sleep(time.Duration(delay) * time.Millisecond)
	rating := rand.Intn(10) + 1
	chnl <- rating
}

func main() {
	chnl := make(chan int)
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go addRating(chnl, &wg)
	}
	go func() {
		wg.Wait()
		close(chnl)
	}()
	res := 0
	cnt := 0
	for x := range chnl {
		res += x
		cnt += 1
	}
	finish := time.Since(start)
	fmt.Println("final rating is", float64(res)/float64(cnt))
	fmt.Println("elapsed time", finish)
}
