package syncObject

import (
	"fmt"
	"runtime"
	"sync"
)

func TestRWMutex() {
	var waitGroup = new(sync.WaitGroup)

	var rwMutex = new(sync.RWMutex)

	var data int = 0

	waitGroup.Add(1)
	go func() {
		defer runtime.Gosched()
		defer waitGroup.Done()
		for i := 0; i < 5; i++ {
			rwMutex.Lock()

			data += 1
			fmt.Println("write 1 goroutine : ", data)

			rwMutex.Unlock()

		}
		runtime.Gosched()
	}()

	waitGroup.Add(1)
	go func() {
		defer runtime.Gosched()
		defer waitGroup.Done()
		for i := 0; i < 5; i++ {
			rwMutex.RLock()

			fmt.Println("read 1 goroutine : ", data)

			rwMutex.RUnlock()
		}
	}()

	waitGroup.Add(1)
	go func() {
		defer runtime.Gosched()
		defer waitGroup.Done()
		for i := 0; i < 5; i++ {
			rwMutex.RLock()

			fmt.Println("read 2 goroutine : ", data)

			rwMutex.RUnlock()
		}
	}()


	waitGroup.Wait()
	fmt.Println("data : ", data)
}
