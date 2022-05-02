package lock

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

//reentrantLock 可重入锁
type reentrantLock struct {
	owner int64
	count int
	lock  *sync.Mutex
}

func newReentrantLock() sync.Locker {
	return &reentrantLock{
		lock: new(sync.Mutex),
	}
}

func (r *reentrantLock) Lock() {
	gid := getGid()
	if atomic.LoadInt64(&r.owner) == gid {
		r.count++
		return
	}
	r.lock.Lock()
	atomic.StoreInt64(&r.owner, gid)
	r.count = 1
}

func (r *reentrantLock) Unlock() {
	gid := getGid()
	if r.count == 0 || atomic.LoadInt64(&r.owner) != gid {
		panic("状态错误")
	}
	r.count--
	if r.count == 0 {
		atomic.StoreInt64(&r.owner, 0)
		r.lock.Unlock()
	}
}

//getGid 通过stack获取goroutine id
func getGid() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(string(buf[:n]))[1]
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		fmt.Println("获取gid失败", err)
	}
	return id
}
