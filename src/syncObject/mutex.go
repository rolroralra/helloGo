package syncObject

import (
	"fmt"
	"runtime"
	"sync"
)

func TestMutex() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := new(sync.WaitGroup)

	var data []int
	//var data = []int{}

	var mutex = new(sync.Mutex)

	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()

			data = append(data, 1)

			mutex.Unlock()
		}

		wg.Done()
		runtime.Gosched()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()

			data = append(data, 1)

			mutex.Unlock()
		}

		wg.Done()
		runtime.Gosched()
	}()

	wg.Wait()
	fmt.Println(len(data))

}
