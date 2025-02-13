package main

import "fmt"

type hungrySingleton struct{}

func (h *hungrySingleton) Call() {
	fmt.Println("Hungry Singleton method called")
}

var hungryInstance = new(hungrySingleton)

func GetHungryInstance() *hungrySingleton {
	return hungryInstance
}
