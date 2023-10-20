package main

import (
	"fmt"
	"sort"
	"sync"
)

func getFreq(str string, chnl chan map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	m := make(map[string]int)
	for _, x := range str {
		m[string(x)]++
	}
	chnl <- m
}

func main() {
	arr := []string{"apple", "trapple", "lol", "cream", "pie"}
	chnl := make(chan map[string]int)
	var wg sync.WaitGroup

	for _, x := range arr {
		wg.Add(1)
		go getFreq(x, chnl, &wg)
	}
	go func() {
		wg.Wait()
		close(chnl)
	}()
	res := make(map[string]int)
	for m := range chnl {
		for k, v := range m {
			res[string(k)] += v
		}
	}
	tmp := make([]string, 0)
	for k, _ := range res {
		tmp = append(tmp, k)
	}
	sort.Strings(tmp)

	for _, x := range tmp {
		fmt.Println("freq of ", x, ": ", res[x])
	}

}
