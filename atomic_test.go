package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				//x += 1 ganti dengan atomic
				atomic.AddInt64(&x, 1)
			}
			defer group.Done()
		}()
	}
	group.Wait()
	fmt.Print("Counter : ", x)
}
