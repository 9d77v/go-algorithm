package lock

import (
	"fmt"
	"sync"
	"testing"
)

func TestReentrantLock(t *testing.T) {
	group := new(sync.WaitGroup)
	var lock = newReentrantLock()
	var k int = 0
	f := func(group *sync.WaitGroup) {
		lock.Lock()
		defer lock.Unlock()
		k++
		fmt.Print("f", k, " ")
	}
	g := func(group *sync.WaitGroup) {
		lock.Lock()
		defer lock.Unlock()
		defer group.Done()
		f(group)
		k++
		fmt.Print("g", k, " ")
	}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go g(group)
	}
	group.Wait()
	fmt.Println()
}
