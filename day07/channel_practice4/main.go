package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type job struct {
	value int64
}

type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100) 
// randomInt()函数的goroutine不停往jobChan中发送随机数，sumInt()函数的24个goroutine不停从jobChan中接收随机数，所以100的容量不会装满，不会出现阻塞
var resultChan = make(chan *result, 100)
// sumInt()函数的24个goroutine不停往resultChan中发送各位数之和，而main函数主goroutine的for range循环从resultChan中不停接收

func randomInt(job1 chan<- *job) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Int63()
		newJob :=&job{
			value: x,
		}
		job1 <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func sumInt(job1 <-chan *job, resChan chan<- *result) {
	defer wg.Done()
	for {
		job := <- job1
		sum := int64(0)
		n := job. value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resChan <- newResult
	}
}

func main() {
	wg.Add(1)
	go randomInt(jobChan)
	for i := 1; i <= 24; i++ {
		wg.Add(1)
		go sumInt(jobChan, resultChan)
	}
	

	for result := range resultChan{
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()

}
