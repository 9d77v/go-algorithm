package lock

import (
	"sync"
	"sync/atomic"
)

//spinLock 自旋锁
type spinLock uint32

func newSpinLock() sync.Locker {
	return new(spinLock)
}

//Lock 加锁
func (s *spinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(s), 0, 1) {
	}
}

//Unlock 解锁
func (s *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(s), 0)
}
