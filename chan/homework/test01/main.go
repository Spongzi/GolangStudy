package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 启动一个协程，将1-2000的数字放入到一个channel中，比如numChan
// 启动8个协程，从numChan取出数(比如n)，计算1+...+n,并存放在resChan
// 最后8个协程同步完成工作后，再遍历resChan,显式结果[如res[1]=1,...res[10]=55]
// 注意：考虑resChan chan int 是否合适

func Sum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func PutNumChan(numChan chan int) {
	defer wg.Done()
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func GetResChan(numChan chan int, resChan chan int) {
	defer wg.Done()
	for i := range numChan {
		res := Sum(i)
		resChan <- res
	}
}

func main() {
	start := time.Now()
	numChan := make(chan int, 2000)
	resChan := make(chan int, 2000)
	wg.Add(1)
	go PutNumChan(numChan)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go GetResChan(numChan, resChan)
	}
	wg.Wait()
	close(resChan)
	lenResChan := len(resChan)
	for i := 1; i < lenResChan; i++ {
		fmt.Printf("map[%d]=%d\n", i, <-resChan)
	}
	end := time.Now()
	timeout := end.Sub(start)
	fmt.Println("一共执行了", timeout)
}
