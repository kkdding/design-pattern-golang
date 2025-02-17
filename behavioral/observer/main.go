package main

import "fmt"

// ----- 抽象层 -----

// Listener 观察者抽象类
type Listener interface {
	FixBug() //观察者发现bug后要做的事
}

type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

// ----- 实现层 -----

// ServiceA 观察者服务
type ServiceA struct {
	State string
}

func (s *ServiceA) FixBug() {
	fmt.Println("服务A修复bug")
}

func (s *ServiceA) CurrentState() {
	fmt.Println("服务A当前状态:", s.State)
}

// ServiceB 观察者服务
type ServiceB struct {
	State string
}

func (s *ServiceB) FixBug() {
	fmt.Println("服务B修复bug")
}

func (s *ServiceB) CurrentState() {
	fmt.Println("服务B当前状态:", s.State)
}

// ServiceC 观察者服务
type ServiceC struct {
	State string
}

func (s *ServiceC) FixBug() {
	fmt.Println("服务C修复bug")
}

func (s *ServiceC) CurrentState() {
	fmt.Println("服务C当前状态:", s.State)
}

// ServiceMonitor 通知者监控
type ServiceMonitor struct {
	listenerList []Listener //需要通知的全部观察者集合
}

func (s *ServiceMonitor) AddListener(listener Listener) {
	s.listenerList = append(s.listenerList, listener)
}

func (s *ServiceMonitor) RemoveListener(listener Listener) {
	for index, l := range s.listenerList {
		//找到要删除的元素位置
		if listener == l {
			//将删除的点前后的元素链接起来
			s.listenerList = append(s.listenerList[:index], s.listenerList[index+1:]...)
			break
		}
	}
}

func (s *ServiceMonitor) Notify() {
	for _, listener := range s.listenerList {
		//依次调用全部观察的具体动作
		listener.FixBug()
	}
}

// 主函数 - 业务逻辑
func main() {
	sa := &ServiceA{
		State: "[1]",
	}
	sb := &ServiceB{
		State: "[2]",
	}
	sc := &ServiceC{
		State: "[3]",
	}

	serverMonitor := new(ServiceMonitor)

	fmt.Println("未出现BUG...")
	sa.CurrentState()
	sb.CurrentState()
	sc.CurrentState()

	serverMonitor.AddListener(sa)
	serverMonitor.AddListener(sb)
	serverMonitor.AddListener(sc)

	fmt.Println("出现BUG...")
	serverMonitor.Notify()
}
