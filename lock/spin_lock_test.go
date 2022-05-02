package lock

import (
	"fmt"
	"sync"
	"testing"
)

func TestSpinLock(t *testing.T) {
	group := new(sync.WaitGroup)
	var lock = newSpinLock()
	var k int = 0
	f := func(group *sync.WaitGroup) {
		lock.Lock()
		defer lock.Unlock()
		defer group.Done()
		k++
		fmt.Print(k, " ")
	}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go f(group)
	}
	group.Wait()
	fmt.Println()
}
