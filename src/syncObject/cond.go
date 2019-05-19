package syncObject

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func TestCond() {
	wg := new(sync.WaitGroup)

	cond := new(sync.Cond)
	cond.L = new(sync.Mutex)

	c := make(chan string, 3)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer runtime.Gosched()
			defer wg.Done()

			cond.L.Lock()

			c <- "go routine " + strconv.Itoa(n) + " : channel written!"

			fmt.Println("wait begin : ", n)

			cond.Wait()

			fmt.Println("wait end : ", n)

			cond.L.Unlock()
		}(i)
	}

	wg2 := new(sync.WaitGroup)
	wg2.Add(1)
	go func() {
		defer runtime.Gosched()
		defer wg2.Done()

		for i := 0; i < 3; i++ {
			fmt.Println("channel value print : ", <-c)
		}
	}()
	wg2.Wait()

	//for i := 0; i < 3; i++ {
	//	fmt.Println("channel value print : ", <-c)
	//}

	for i := 0; i < 3; i++ {
		cond.L.Lock()

		fmt.Println("signal : ", i)
		cond.Signal()

		cond.L.Unlock()
	}

	wg.Wait()
}
