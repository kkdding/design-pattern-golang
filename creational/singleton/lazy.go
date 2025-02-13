package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var initialized uint32
var lock sync.Mutex

type lazySingleton struct{}

func (l *lazySingleton) Call() {
	fmt.Println("Lazy Singleton method called")
}

var lazyInstance *lazySingleton

func GetLazyInstance() *lazySingleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return lazyInstance
	}

	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		lazyInstance = new(lazySingleton)
		atomic.StoreUint32(&initialized, 1)
	}
	return lazyInstance
}
