package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTaskControl(t *testing.T) {
	taskNum := 5
	wg := sync.WaitGroup{}
	wg.Add(taskNum)

	for i := 0; i < taskNum; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("info", i)
		}(i)
	}
	wg.Wait()
}

func TestPrint(t *testing.T) {
	fmt.Println("hello")
}

func TestA(t *testing.T) {
	test := make(chan int, 10)
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func(info chan int) {
		defer wg.Done()
		for {
			select {
			case val, ok := <-test:
				if !ok {
					t.Logf("Channel Closed!")
					return
				}

				t.Logf("data %d\n", val)
			}
		}
	}(test)

	go func() {
		defer wg.Done()
		test <- 1
		time.Sleep(1 * time.Second)
		test <- 2

		close(test)
	}()
	wg.Wait()
}
