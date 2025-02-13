package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

func (s *singleton) Call() {
	fmt.Println("Singleton method called")
}

var once sync.Once

var instance *singleton

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}
