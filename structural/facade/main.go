package main

import "fmt"

// SystemA 实现类
type SystemA struct{}

func (s *SystemA) StartFuncA1() {
	fmt.Println("function-A1 start...")
}

func (s *SystemA) StartFuncA2() {
	fmt.Println("function-A2 start...")
}

// SystemB 实现类
type SystemB struct{}

func (s *SystemB) StartFuncB1() {
	fmt.Println("function-B1 start...")
}

// SystemC 实现类
type SystemC struct{}

func (s *SystemC) StartFuncC1() {
	fmt.Println("function-C1 start...")
}

// Facade 外观类 统筹接口
type Facade struct {
	systemA *SystemA
	systemB *SystemB
	systemC *SystemC
}

func (f *Facade) StartA1B1() {
	f.systemA.StartFuncA1()
	f.systemB.StartFuncB1()
}

func (f *Facade) StartA2C1() {
	f.systemA.StartFuncA2()
	f.systemC.StartFuncC1()
}

func (f *Facade) StartB1C1() {
	f.systemB.StartFuncB1()
	f.systemC.StartFuncC1()
}

// 主函数 - 业务逻辑
func main() {
	facade := &Facade{
		systemA: &SystemA{},
		systemB: &SystemB{},
		systemC: &SystemC{},
	}

	facade.StartA1B1()
	facade.StartA2C1()
	facade.StartB1C1()

}
